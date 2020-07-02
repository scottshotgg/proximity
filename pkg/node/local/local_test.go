package local_test

import (
	"fmt"
	"testing"
	"time"

	_ "net/http/pprof"

	"github.com/paulbellamy/ratecounter"
	"github.com/scottshotgg/proximity/pkg/node"
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

	go recver(l)
	// go sender(l)

	sender(l)
}

func sender(l *local.Local) {
	// var i int

	// go func() {
	// 	for {
	// 		if i == 1000000000 {
	// 			return
	// 		}

	// 		fmt.Println("i:", i)
	// 		fmt.Println("Percent:", math.Round(100*float64(i)/float64(1000000000))/100)

	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	// for ; i < 1000000000; i++ {
	for {
		l.Send(&node.Msg{
			Route: "a",
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
		// select {
		<-ch
		recvcounter.Incr(1)

		// case <-time.After(1 * time.Second):
		// 	return
		// }
	}
}
