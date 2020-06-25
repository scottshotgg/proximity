package local_test

import (
	"fmt"
	"testing"
	"time"

	_ "net/http/pprof"

	"github.com/paulbellamy/ratecounter"
	"github.com/scottshotgg/proximity/pkg/node/local"
	// "go.uber.org/goleak"
)

var recvcounter = ratecounter.NewRateCounter(1 * time.Second)
var sendcounter = ratecounter.NewRateCounter(1 * time.Second)

var contents []byte

func TestP2P(t *testing.T) {
	var l = local.New()

	go func() {
		var timer = time.NewTimer(1 * time.Second)

		for {
			select {
			case <-timer.C:
				fmt.Println("Recv Rate:", recvcounter.Rate())
				fmt.Println("Send Rate:", sendcounter.Rate())

				timer.Reset(1 * time.Second)
			}
		}
	}()

	// go recver(l)
	time.Sleep(100 * time.Millisecond)

	// for i := 0; i < 9; i++ {
	// 	go sender(l)
	// }

	go sender(l)
	recver(l)
}

func sender(l *local.Local) {
	for {
		l.Send(&local.Msg{
			Route:    "a",
			Contents: contents,
		})

		// time.Sleep(200 * time.Millisecond)

		// sendcounter.Incr(1)

		// ch <- &local.Msg{
		// 	Route:    "a",
		// 	Contents: []byte("hey its me"),
		// }

		sendcounter.Incr(1)
	}
}

func recver(l *local.Local) {
	var ch = l.Join("a")

	for {
		select {
		case _ = <-ch:
			recvcounter.Incr(1)
		}
	}
}
