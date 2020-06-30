package client

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

func Start(addr string) {
	serverAddr, err := net.ResolveTCPAddr("tcp", addr)
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
		log.Fatalln("err DialTCP:", err)
	}

	for {
		fmt.Fprintln(conn, strings.Repeat("a", 1*1024*1024))
	}
}
