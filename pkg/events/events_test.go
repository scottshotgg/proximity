package events_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/scottshotgg/proximity/pkg/events"
)

func Test_something(t *testing.T) {
	var (
		e        = events.New()
		sendSize = 10
		recvSize = 1

		recvCounts = make([]int64, recvSize)
		sendCounts = make([]int64, sendSize)
		// recvCount int64
		// sendCount int64
	)

	go func() {
		var timer = time.NewTimer(1 * time.Second)

		for {
			select {
			case <-timer.C:
				var (
					// rc = recvCount
					rc = recvCounts
					// sc = sendCount
					sc = sendCounts
				)

				recvCounts = make([]int64, recvSize)
				sendCounts = make([]int64, sendSize)
				// recvCount = 0
				// sendCount = 0
				var sendTotal int64
				var recvTotal int64

				if recvSize > 1 {
					for _, count := range rc {
						recvTotal += count
					}
					fmt.Println("Recv count:", recvTotal)
				} else {
					fmt.Println("Recv count:", rc)
				}

				if sendSize > 1 {
					for _, count := range sc {
						sendTotal += count
					}

					fmt.Println("Send count:", sendTotal)
				} else {
					fmt.Println("Send count:", sc)
				}

				fmt.Println()

				timer.Reset(1 * time.Second)
			}
		}
	}()

	var wg = &sync.WaitGroup{}

	for i := 0; i < recvSize; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			var ch = e.Listen("a", "a")

			for range ch {
				recvCounts[id]++
				// atomic.AddInt64(&recvCount, 1)
			}
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	for i := 0; i < sendSize; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			var err error

			for {
				err = e.SendMulti([]*events.Msg{
					{
						Route: "a",
					},
				})

				if err != nil {
					fmt.Println("err:", err)
					continue
				}

				sendCounts[id]++
				// atomic.AddInt64(&sendCount, 1)
			}
		}(i)
	}

	wg.Wait()
}
