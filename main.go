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

		b = channel_bus.New(10000000)
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
		max        = 2
	)

	const writersPerReaders = 20

	for i := 0; i < max-1; i++ {
		for j := 0; j < writersPerReaders; j++ {
			fmt.Println("sender")
			go send(j+(i*writersPerReaders), sendClient)
		}

		go recv(i, recvClient)
	}

	for j := 0; j < writersPerReaders; j++ {
		fmt.Println("sender")
		go send(j+(max*writersPerReaders), sendClient)
	}

	recv(max, recvClient)
}

func recv(id int, recvClient buffs.RecvClient) {
	var ctx = context.Background()

	var listener, err = recvClient.Attach(ctx, &buffs.AttachReq{
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

func send(id int, sendClient buffs.SenderClient) {
	var (
		i int

		ctx           = context.Background()
		everySecond   = 1 * time.Second
		counter       = ratecounter.NewRateCounter(everySecond)
		timer         = time.NewTimer(everySecond)
		sendPipe, err = sendClient.Send(ctx)
	)

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
