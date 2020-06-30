package client

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

func Start() {
	serverAddr, err := net.ResolveTCPAddr("tcp", "localhost:9090")
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	var wg = &sync.WaitGroup{}

	wg.Add(1)

	for i := 0; i < 10; i++ {
		go send(serverAddr)
	}

	wg.Wait()
}

func send(serverAddr *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatalln("err ListenTCP:", err)
	}

	for {
		fmt.Fprintln(conn, strings.Repeat("a", 1*1024*1024))
	}
}
