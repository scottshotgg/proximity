package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/scottshotgg/proximity/pkg/buffs"
	grpc_node "github.com/scottshotgg/proximity/pkg/node/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var err = start()
	if err != nil {
		// TODO: need to wrap errors
		log.Fatalln("err start:", err)
	}
}

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func start() error {
	const (
		port        = 5001
		mb          = 1024 * 1024
		everySecond = 1 * time.Second
		maxMsgMB    = 100
		maxMsgSize  = maxMsgMB * mb
	)

	var ip = getOutboundIP()

	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	const maxMsgSizeHeader = "Max message size"
	var mmshLen = strconv.Itoa(len(maxMsgSizeHeader) + 1)

	log.Printf("%-"+mmshLen+"s: %s\n", "Address", ip.String())
	log.Printf("%-"+mmshLen+"s: %d\n", "Port", port)
	log.Printf("%-"+mmshLen+"s: %d MB, %d bytes\n", maxMsgSizeHeader, maxMsgMB, maxMsgSize)
	fmt.Println()

	var (
		grpcServer = grpc.NewServer(
			grpc.MaxSendMsgSize(maxMsgSize),
			grpc.MaxRecvMsgSize(maxMsgSize),
		)
	)

	go ctrlc(grpcServer)

	buffs.RegisterNodeServer(grpcServer, grpc_node.New())
	log.Printf("Registered node")

	reflection.Register(grpcServer)
	log.Printf("Registered reflection")

	log.Println("Serving gRPC ...")
	fmt.Println()

	return grpcServer.Serve(l)
}

func ctrlc(s *grpc.Server) {
	var c = make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Stopping gRPC server ...")

	// TODO: flesh this out a bit more
	s.GracefulStop()
}
