package events_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/scottshotgg/proximity/pkg/events"
)

func Test_something(t *testing.T) {
	var (
		// sendcounter = ratecounter.NewRateCounter(1 * time.Second)
		// ratecounter = ratecounter.NewRateCounter(1 * time.Second)

		e = events.New()

		recvCount int64
		sendCount int64

		// Register top level listener
		ev, _ = e.Listen("a", "a")
	)

	// Register ID listener
	var evv, _ = ev.Listen("something_else", "something_else")
	var _, chhh1 = evv.Listen("blah", "blah")
	var _, chhh2 = evv.Listen("blurb", "blurb")

	go func() {
		var timer = time.NewTimer(1 * time.Second)

		for {
			select {
			case <-timer.C:
				var (
					rc = recvCount
					sc = sendCount
				)

				recvCount = 0
				sendCount = 0

				fmt.Println("Recv count:", sc)
				fmt.Println("Send count:", rc)
				fmt.Println()

				timer.Reset(1 * time.Second)
			}
		}
	}()

	for i := 0; i < 1; i++ {
		go func(id int) {
			for range chhh1 {
				// fmt.Println("got message for blah")
				recvCount++
			}

			// for range chh {
			// 	// atomic.AddInt64(&recvCount, 1)
			// }
		}(i)

		go func(id int) {
			for range chhh2 {
				// fmt.Println("got message for blurb")
				recvCount++
			}

			// for range chh {
			// recvCount++
			// 	// atomic.AddInt64(&recvCount, 1)
			// }
		}(i)
	}

	for {
		var err = e.Handle("a/something_else/", &events.Msg{
			Route: "a/something_else/",
		})

		if err == nil {
			sendCount++
		}

		// time.Sleep(100 * time.Millisecond)

		// atomic.AddInt64(&sendCount, 1)
	}
}
