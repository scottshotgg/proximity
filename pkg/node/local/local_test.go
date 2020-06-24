package local_test

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	_ "net/http/pprof"

	"github.com/paulbellamy/ratecounter"
	"github.com/scottshotgg/proximity/pkg/node/local"

	// "go.uber.org/goleak"
	"github.com/pkg/profile"
)

func TestP2P(t *testing.T) {
	defer profile.Start().Stop()
	// go http.ListenAndServe("0.0.0.0:8080", nil)

	// defer goleak.VerifyNone(t)

	const (
		route   = "a"
		testMsg = ""
	)

	var (
		n        = local.New()
		pub, err = n.Publish(route)
	)

	if err != nil {
		log.Fatalln("n.Publish")
	}

	// Subscribe to the route
	sub, _, err := n.Subscribe(route)
	if err != nil {
		log.Fatalln("n.Subscribe")
	}

	const everySecond = 1 * time.Second

	var (
		recvcounter = ratecounter.NewRateCounter(everySecond)
		sendcounter = ratecounter.NewRateCounter(everySecond)
		timer       = time.NewTimer(everySecond)
		testTime    = 20 * everySecond
	)

	var ctx, cancel = context.WithTimeout(context.Background(), testTime)
	defer cancel()

	go func() {
		for {
			select {
			case <-timer.C:
				log.Println("Recv:", recvcounter.Rate())
				log.Println("Send:", sendcounter.Rate())
				timer.Reset(everySecond)

			case <-ctx.Done():
				timer.Stop()
				return
			}
		}
	}()

	// go func() {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		}
	// 	}
	// }()

	var wg = &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return

			// Insert test message
			case pub <- []byte{}:
				sendcounter.Incr(1)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return

			case <-sub:
				recvcounter.Incr(1)
			}
		}
	}()

	wg.Wait()
	cancel()

	// close(pub)
	n.Close()
	close(sub)
}
