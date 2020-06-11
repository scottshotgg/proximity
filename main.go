package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/paulbellamy/ratecounter"
	buffs "github.com/scottshotgg/proximity/pkg/buffs"
	bus "github.com/scottshotgg/proximity/pkg/bus"
	"github.com/scottshotgg/proximity/pkg/node"
	channel_recv "github.com/scottshotgg/proximity/pkg/recv/channel"
	grpc_recv "github.com/scottshotgg/proximity/pkg/recv/grpc"
	channel_sender "github.com/scottshotgg/proximity/pkg/sender/channel"
	grpc_sender "github.com/scottshotgg/proximity/pkg/sender/grpc"
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
	var n = node.New()

	go n.Start(5001)

	time.Sleep(100 * time.Millisecond)

	// time.AfterFunc(1*time.Second, func() {
	// 	os.Exit(0)
	// })

	go recv(0, n)
	send(0, n)
}

func recv(id int, n *node.Node) {
	var conn, err = grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err recvConn:", err)
	}

	var (
		recvClient = buffs.NewNodeClient(conn)
		ctx        = context.Background()
	)

	listener, err := recvClient.Attach(ctx, &buffs.AttachReq{
		Id:    strconv.Itoa(id),
		Route: "a",
	})

	if err != nil {
		log.Fatalln("err making listener:", err)
	}

	md, err := listener.Header()
	if err != nil {
		log.Fatalln("err getting header values:", err)
	}

	fmt.Println("metadata:", md)

	var counter = ratecounter.NewRateCounter(1 * time.Second)

	var timer = time.NewTimer(1 * time.Second)

	go func() {
		for {
			var _, err = listener.Recv()
			if err != nil {
				log.Fatalln("err", err)
			}

			counter.Incr(1)
		}
	}()

	for {
		select {
		case <-timer.C:
			fmt.Printf("Rate for %d: %d\n", id, counter.Rate())
			timer.Reset(1 * time.Second)
		}
	}
}

func send(id int, n *node.Node) {
	var conn, err = grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err recvConn:", err)
	}

	var (
		sendClient = buffs.NewNodeClient(conn)

		i int

		ctx         = context.Background()
		everySecond = 1 * time.Second
		counter     = ratecounter.NewRateCounter(everySecond)
		timer       = time.NewTimer(everySecond)
	)

	sendPipe, err := sendClient.Send(ctx)
	if err != nil {
		log.Fatalln("err Send:", err)
	}

	go func() {
		for {
			err = sendPipe.Send(&buffs.SendReq{
				Msg: &buffs.Message{
					Route:    "a",
					Contents: "",
				},
			})

			if err != nil {
				if err != io.EOF {
					log.Fatalln("err sending:", err)
				}
			}

			i++
			counter.Incr(1)

			// time.Sleep(50 * time.Millisecond)
		}
	}()

	for {
		select {
		case <-timer.C:
			// fmt.Printf("Send rate for %d: %d\n", id, counter.Rate())

			timer.Reset(everySecond)
		}
	}
}

func servers(b bus.Bus) {
	go func() {
		var err = grpc_sender.New(5001, channel_sender.New(b))
		if err != nil {
			log.Fatalln("err created sender:", err)
		}
	}()

	go func() {
		var err = grpc_recv.New(5002, channel_recv.New(b))
		if err != nil {
			log.Fatalln("err creating recv", err)
		}
	}()
}
