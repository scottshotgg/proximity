package main

import (
	"log"
	"net"
	"time"

	"github.com/scottshotgg/proximity/pkg/buffs"
	grpc_node "github.com/scottshotgg/proximity/pkg/node/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var l, err = net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalln("err net.Listen:", err)
	}

	const (
		mb          = 1024 * 1024
		everySecond = 1 * time.Second
		maxMsgSize  = 100 * mb
	)

	var (
		grpcServer = grpc.NewServer(
			grpc.MaxSendMsgSize(maxMsgSize),
			grpc.MaxRecvMsgSize(maxMsgSize),
		)
	)

	buffs.RegisterNodeServer(grpcServer, grpc_node.New())
	reflection.Register(grpcServer)

	grpcServer.Serve(l)
}
