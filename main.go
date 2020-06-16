package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/paulbellamy/ratecounter"
	buffs "github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/listener"
	grpc_lis "github.com/scottshotgg/proximity/pkg/listener/grpc"
	grpc_node "github.com/scottshotgg/proximity/pkg/node/grpc"
	"google.golang.org/grpc"
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

	var n = servers(port)

	time.Sleep(100 * time.Millisecond)

	clients(n)

	// time.AfterFunc(1*time.Second, func() {
	// 	os.Exit(0)
	// })
}

func clients(n *grpc_node.Node) {
	const port = ":5001"

	var wg = &sync.WaitGroup{}

	var funcs = []func(id int, c buffs.NodeClient){
		send,
	}

	for i := range funcs {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			var conn, err = grpc.Dial(port, grpc.WithInsecure())
			if err != nil {
				log.Fatalln("err recvConn:", err)
			}

			funcs[id](id, buffs.NewNodeClient(conn))
		}(i)
	}

	go recv(1)

	wg.Wait()
}

func recv(id int) {

	// TODO: add checking for id and route

	var (
		counter = ratecounter.NewRateCounter(1 * time.Second)
		_, err  = grpc_lis.New(strconv.Itoa(id), "a", ":5001", func(msg *listener.Msg) error {
			counter.Incr(1)

			return nil
		})
	)

	if err != nil {
		log.Fatalln("error creating listener:", err)
	}

	var timer = time.NewTimer(1 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Printf("Rate for %d: %d\n", id, counter.Rate())

			timer.Reset(1 * time.Second)
		}
	}
}

func send(id int, c buffs.NodeClient) {
	var (
		i int

		ctx         = context.Background()
		everySecond = 1 * time.Second
		counter     = ratecounter.NewRateCounter(everySecond)
		timer       = time.NewTimer(everySecond)
		emptyMsg    = &buffs.Message{
			Route:    "a",
			Contents: "",
		}

		sendPipe, err = c.Send(ctx)
	)

	if err != nil {
		log.Fatalln("err Send:", err)
	}

	go func() {
		for {
			select {
			case <-timer.C:
				// fmt.Printf("Send rate for %d: %d\n", id, counter.Rate())

				timer.Reset(everySecond)
			}
		}
	}()

	for {
		err = sendPipe.Send(&buffs.SendReq{
			Msg: emptyMsg,
		})

		if err != nil {
			log.Fatalln("err sending:", err)
		}

		i++
		counter.Incr(1)
	}
}

func servers(port int) *grpc_node.Node {
	var n = grpc_node.New()

	go n.Start(port)

	return n
}
