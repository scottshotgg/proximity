package local_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/paulbellamy/ratecounter"
	"github.com/scottshotgg/proximity/pkg/node/local"
)

func TestP2P(t *testing.T) {
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
	sub, err := n.Subscribe(route)
	if err != nil {
		log.Fatalln("n.Subscribe")
	}

	const everySecond = 1 * time.Second

	var (
		recvcounter = ratecounter.NewRateCounter(everySecond)
		sendcounter = ratecounter.NewRateCounter(everySecond)
		timer       = time.NewTimer(everySecond)
	)

	time.AfterFunc(1*time.Minute, func() {
		os.Exit(9)
	})

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Recv:", recvcounter.Rate())
				fmt.Println("Send:", sendcounter.Rate())
				timer.Reset(everySecond)
			}
		}
	}()

	go func() {
		for {
			// Insert test message
			pub <- []byte{}
			sendcounter.Incr(1)
		}
	}()

	for {
		select {
		case <-sub:
			recvcounter.Incr(1)
		}
	}
}
