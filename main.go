package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/paulbellamy/ratecounter"
	buffs "github.com/scottshotgg/proximity/pkg/buffs"
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
	var wg = &sync.WaitGroup{}

	var conn, err = grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err recvConn:", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		recv("0", buffs.NewNodeClient(conn))
	}()

	conn, err = grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err recvConn:", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		send(1, buffs.NewNodeClient(conn))
	}()

	wg.Wait()
}

func recv(id string, c buffs.NodeClient) {
	var (
		ctx   = context.Background()
		route = "a"

		listener, err = c.Attach(ctx, &buffs.AttachReq{
			Id:    id,
			Route: route,
		})
	)

	if err != nil {
		log.Fatalln("err making listener:", err)
	}

	// TODO: metadata not implemented right now for Node
	md, err := listener.Header()
	if err != nil {
		log.Fatalln("err getting header values:", err)
	}

	var (
		idHeader    = md.Get("id")
		routeHeader = md.Get("route")
	)

	if len(idHeader) > 0 {
		id = idHeader[0]
	}

	if len(routeHeader) > 0 {
		route = routeHeader[0]
	}

	// TODO: add checking for id and route

	var counter = ratecounter.NewRateCounter(1 * time.Second)

	go func() {
		var timer = time.NewTimer(1 * time.Second)

		for {
			select {
			case <-timer.C:
				fmt.Printf("Rate for %s: %d\n", id, counter.Rate())

				timer.Reset(1 * time.Second)
			}
		}
	}()

	for {
		var _, err = listener.Recv()
		if err != nil {
			log.Fatalln("err", err)
		}

		counter.Incr(1)
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
