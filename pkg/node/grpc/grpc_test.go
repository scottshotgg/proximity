package grpc_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/paulbellamy/ratecounter"
	"github.com/scottshotgg/proximity/pkg/buffs"
	grpcNode "github.com/scottshotgg/proximity/pkg/node/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func TestP2P(t *testing.T) {
	const (
		mb          = 1024 * 1024
		everySecond = 1 * time.Second

		key  = "/home/scottshotgg/Development/go/me/proximity/server.key"
		cert = "/home/scottshotgg/Development/go/me/proximity/server.crt"
	)

	var (
		maxMsgSize = 100 * mb

		l, err = net.Listen("tcp", "localhost:5001")
	)

	if err != nil {
		log.Fatalf("could not load TLS keys: %s\n", err)
	}

	// // Create the TLS credentials
	// serverCreds, err := credentials.NewServerTLSFromFile(cert, key)
	// if err != nil {
	// 	log.Fatalf("could not load TLS keys: %s\n", err)
	// }

	var grpcServer = grpc.NewServer(
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.MaxRecvMsgSize(maxMsgSize),
		// grpc.Creds(serverCreds),
	)

	buffs.RegisterNodeServer(grpcServer, grpcNode.New())
	reflection.Register(grpcServer)

	go grpcServer.Serve(l)

	var (
		counter     = ratecounter.NewRateCounter(everySecond)
		timer       = time.NewTimer(everySecond)
		sendcounter = ratecounter.NewRateCounter(everySecond)
	)

	time.AfterFunc(1*time.Minute, func() {
		os.Exit(9)
	})

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Rate:", counter.Rate())
				fmt.Println("Send:", sendcounter.Rate())

				timer.Reset(everySecond)
			}
		}
	}()

	var defaultOps = grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(maxMsgSize),
		grpc.MaxCallSendMsgSize(maxMsgSize),
	)

	var data []byte
	var routes = []string{"a"}

	// // Create the client TLS credentials
	// clientCreds, err := credentials.NewClientTLSFromFile(cert, "")
	// if err != nil {
	// 	log.Fatalf("could not load tls cert: %s", err)
	// }

	for i := 0; i < 2; i++ {
		conn, err := grpc.Dial("localhost:5001",
			grpc.WithInsecure(),
			defaultOps,
			// grpc.WithTransportCredentials(clientCreds),
		)
		if err != nil {
			t.Errorf("%v", err)
		}

		var client = buffs.NewNodeClient(conn)

		pub, err := client.Publish(context.Background())
		if err != nil {
			log.Fatalln("err", err)
		}

		for {
			err = pub.Send(&buffs.PublishReq{
				Routes:   routes,
				Contents: data,
			})

			if err != nil {
				t.Errorf("pub %v", err)
			}

			sendcounter.Incr(1)
		}
	}

	time.Sleep(10 * time.Second)

	// time.Sleep(2000 * time.Millisecond)

	// conn2, err := grpc.Dial(":5001",
	// 	grpc.WithInsecure(),
	// 	defaultOps,
	// 	grpc.WithTransportCredentials(clientCreds),
	// )
	// if err != nil {
	// 	t.Errorf("%v", err)
	// }

	// var client2 = buffs.NewNodeClient(conn2)

	// sub2, err := client2.Subscribe(context.Background(), &buffs.SubscribeReq{
	// 	Topics: []string{"a"},
	// })

	// if err != nil {
	// 	t.Errorf("err1 %v", err)
	// }

	// for {
	// 	_, err = sub2.Recv()
	// 	if err != nil {
	// 		t.Errorf("sub %v", err)
	// 		t.FailNow()
	// 	}

	// 	counter.Incr(1)
	// }
}
