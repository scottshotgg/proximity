package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/scottshotgg/proximity/pkg/buffs"
	grpc_node "github.com/scottshotgg/proximity/pkg/node/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("\n%+v\n", m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Println()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	// defer profile.Start(profile.GoroutineProfile, profile.ProfilePath("profile.p")).Stop()

	// go func() {
	// 	for {
	// 		PrintMemUsage()
	// 		runtime.GC()

	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	start()
	// var err = start()
	// if err != nil {
	// 	// TODO: need to wrap errors
	// 	log.Fatalln("err start:", err)
	// }
}

// Get preferred outbound ip of this machine
func getOutboundIP(host, port string) (net.IP, error) {
	if host == "" {
		host = "8.8.8.8"
		port = "80"
	}

	conn, err := net.Dial("udp", net.JoinHostPort(host, port))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	localAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return nil, errors.New("localAddr was not a *net.UDPAddr")
	}

	return localAddr.IP, nil
}

func start() error {
	const (
		host        = ""
		port        = "5001"
		mb          = 1024 * 1024
		everySecond = 1 * time.Second
		maxMsgMB    = 100
		maxMsgSize  = maxMsgMB * mb
	)

	var (
		ip, err = getOutboundIP(host, port)
	)

	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	const maxMsgSizeHeader = "Max message size"
	var mmshLen = strconv.Itoa(len(maxMsgSizeHeader) + 1)

	log.Printf("%-"+mmshLen+"s: %s\n", "Address", ip.String())
	log.Printf("%-"+mmshLen+"s: %s\n", "Port", port)
	log.Printf("%-"+mmshLen+"s: %d MB, %d bytes\n", maxMsgSizeHeader, maxMsgMB, maxMsgSize)
	fmt.Println()

	var grpcServer = grpc.NewServer(
		grpc.MaxSendMsgSize(maxMsgSize),
		grpc.MaxRecvMsgSize(maxMsgSize),
	)

	go ctrlc(grpcServer)

	buffs.RegisterNodeServer(grpcServer, grpc_node.New())
	log.Printf("Registered node")

	reflection.Register(grpcServer)
	log.Printf("Registered reflection")

	log.Println("Serving gRPC ...")
	fmt.Println()

	// go func() {
	// 	// time.Sleep(12 * time.Second)

	// 	grpcServer.Stop()
	// }()

	return grpcServer.Serve(l)
}

func ctrlc(s *grpc.Server) {
	var c = make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Stopping gRPC server ...")

	// TODO: flesh this out a bit more
	s.Stop()
}
