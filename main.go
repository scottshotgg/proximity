package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/paulbellamy/ratecounter"
	buffs "github.com/scottshotgg/proximity/pkg/buffs"
	bus "github.com/scottshotgg/proximity/pkg/bus"
	channel_bus "github.com/scottshotgg/proximity/pkg/bus/channel"
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
	var (
		// route1 = "ur_mom"
		// route2 = "ur_dad"

		b = channel_bus.New(100)
	)

	go servers(b)

	time.Sleep(1 * time.Second)

	var recvConn, err = grpc.Dial(":5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err recvConn:", err)
	}

	sendConn, err := grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err sendConn:", err)
	}

	var (
		sendClient = buffs.NewSenderClient(sendConn)
		recvClient = buffs.NewRecvClient(recvConn)
		ctx        = context.Background()
	)

	listener, err := recvClient.Attach(ctx, &buffs.AttachReq{
		Id:    "1",
		Route: "a",
	})

	if err != nil {
		log.Fatalln("err making listener:", err)
	}

	go send(sendClient)

	recv(listener)
}

func recv(listener buffs.Recv_AttachClient) {
	md, err := listener.Header()
	if err != nil {
		log.Fatalln("err getting header values:", err)
	}

	fmt.Println("metadata:", md)

	var counter = ratecounter.NewRateCounter(1 * time.Second)

	var timer = time.NewTimer(1 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("Rate:", counter.Rate())
			timer.Reset(1 * time.Second)

		default:
		}

		_, err = listener.Recv()
		if err != nil {
			log.Fatalln("err listening:", err)
		}

		counter.Incr(1)
	}
}

func send(sendClient buffs.SenderClient) {
	var (
		i   int
		ctx = context.Background()
		err error
	)

	for {
		_, err = sendClient.Send(ctx, &buffs.SendReq{
			Msg: &buffs.Message{
				Route:    "a",
				Contents: "else:::::::" + strconv.Itoa(i),
			},
		})

		if err != nil {
			log.Fatalln("err sending:", err)
		}

		i++

		// time.Sleep(50 * time.Millisecond)
	}
}

func servers(b bus.Bus) {
	go func() {
		var err = grpc_recv.New(5002, channel_recv.New(b))
		if err != nil {
			log.Fatalln("err creating recv", err)
		}
	}()

	go func() {
		var err = grpc_sender.New(5001, channel_sender.New(b))
		if err != nil {
			log.Fatalln("err created sender:", err)
		}
	}()
}
