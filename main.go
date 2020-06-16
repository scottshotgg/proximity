package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/paulbellamy/ratecounter"
	grpc_lis "github.com/scottshotgg/proximity-go/listener/grpc"
	grpc_sender "github.com/scottshotgg/proximity-go/sender/grpc"
	"github.com/scottshotgg/proximity/pkg/listener"
	grpc_node "github.com/scottshotgg/proximity/pkg/node/grpc"
)

// func main() {
// 	var (
// 		err error
// 		msg string

// 		c = channel.New(100)
// 	)

// 	// for i := 0; i < 1000; i++ {
// 	err = c.Insert("something here")
// 	if err != nil {
// 		log.Fatalln("err:", err)
// 	}

// 	fmt.Println("inserted")
// 	// }

// 	msg, err = c.Remove()
// 	if err != nil {
// 		log.Fatalln("err:", err)
// 	}

// 	log.Println("msg:", msg)
// }

func main() {
	const port = 5001

	servers(port)

	time.Sleep(100 * time.Millisecond)

	clients()
}

func clients() {
	var (
		wg = &sync.WaitGroup{}

		funcs = []func(id int){
			// send,
			recv,
		}
	)

	for i := range funcs {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			funcs[id](id)
		}(i)
	}

	wg.Wait()
}

func listenerCallback(counter *ratecounter.RateCounter) func(msg *listener.Msg) error {
	return func(msg *listener.Msg) error {
		counter.Incr(1)

		return nil
	}
}

func recv(id int) {
	const (
		addr        = ":5001"
		route       = "a"
		everySecond = 1 * time.Second
		rateFor     = "Rate for %d: %d\n"
	)

	var (
		counter = ratecounter.NewRateCounter(everySecond)
		timer   = time.NewTimer(everySecond)
		_, err  = grpc_lis.New(strconv.Itoa(id), route, addr, listenerCallback(counter))
	)

	if err != nil {
		log.Fatalln("error creating listener:", err)
	}

	// Loop forever and print the rate
	for {
		select {
		case <-timer.C:
			fmt.Printf(rateFor, id, counter.Rate())

			timer.Reset(everySecond)
		}
	}
}

func send(id int) {
	const (
		route       = "a"
		contents    = ""
		everySecond = 1 * time.Second
		addr        = ":5001"
	)

	var (
		blank = []byte(strings.Repeat("a", 0))

		counter = ratecounter.NewRateCounter(everySecond)
		timer   = time.NewTimer(everySecond)

		s, err = grpc_sender.New(strconv.Itoa(id), route, addr)
	)

	if err != nil {
		log.Fatalln("err grpc_sender.New:", err)
	}

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Printf("Send rate for %d: %d\n", id, counter.Rate())

				timer.Reset(everySecond)
			}
		}
	}()

	var t = time.After(1 * time.Minute)

	for {
		select {
		case <-t:
			return

		default:
		}

		err = s.Send(blank)
		if err != nil {
			log.Fatalln("err s.Send:", err)
		}

		counter.Incr(1)
	}
}

func servers(port int) *grpc_node.Node {
	var n = grpc_node.New()

	go n.Start(port)

	return n
}
