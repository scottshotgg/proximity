package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/inhies/go-bytesize"
)

var countBytes int64
var count int64
var total int64
var totalBytes int64

func Start(addr string) {
	serverAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	var start = time.Now()

	go func() {
		<-sigChan

		totalBytes += countBytes
		total += count

		var (
			elapsed = time.Now().Sub(start)
			avg     = float64(totalBytes) / elapsed.Seconds()
		)

		fmt.Println("Stats:")
		fmt.Println("Avg:", bytesize.New(avg))

		os.Exit(9)
	}()

	var ticker = time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			fmt.Printf("Send Count: %v\n", count)
			fmt.Printf("Send Bytes: %v\n", bytesize.New(float64(countBytes)))
			totalBytes += countBytes
			total += count
			countBytes = 0
			count = 0
		}
	}()

	var wg = &sync.WaitGroup{}

	wg.Add(1)

	for i := 0; i < 10; i++ {
		go send(serverAddr)
	}

	wg.Wait()
}

const (
	B  = 1
	KB = 1024 * B

	amount = 64 * KB
)

func send(serverAddr *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatalln("err DialTCP:", err)
	}

	// Set the TCP window to 64KB
	conn.SetWriteBuffer(amount)

	var (
		line int

		br   = bufio.NewWriter(conn)
		a    = strings.Repeat("a", 1) + ":"
		d    = strings.Repeat(a, KB-1)
		data = append([]byte(d), '\n')
	)

	for {
		// Write the delimited messages to the buffer
		line, err = br.Write(data)
		if err != nil {
			log.Fatalln("err fmt.Fprintf:", err)
		}

		// Check if the amount sent is the same as the length of the data
		if line != len(data) {
			log.Fatalln("wtf dude", line, len(data))
		}

		atomic.AddInt64(&countBytes, int64(line))
		atomic.AddInt64(&count, 1)
	}
}
