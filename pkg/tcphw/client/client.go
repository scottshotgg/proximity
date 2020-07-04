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

func Start(addr string, times int) {
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

	for i := 0; i < times; i++ {
		go send(addr)
	}

	wg.Wait()
}

const (
	B  = 1
	KB = 1024 * B

	amount = 131072
)

func send(addr string) {
	serverAddr, err := net.ResolveTCPAddr("tcp", addr+":9090")
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	// var timer = time.NewTimer(5 * time.Second)

	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatalln("err DialTCP:", err)
	}

	// Set the TCP window to 64KB
	conn.SetWriteBuffer(amount)
	conn.SetReadBuffer(amount)

	var (
		line int

		br = bufio.NewWriter(conn)
	)

	// Sender
	err = br.WriteByte('1')
	if err != nil {
		log.Fatalln("br.WriteByte():", err)
	}

	var (
		// Smallest:
		// size = 1

		// Tiny:
		// size = 5

		// Smaller:
		// size = 10

		// Small:
		size = 100

		// Medium:
		// size = 400

		// Below average TCP:
		// size = 1000

		// Average TCP:
		// size = 1500

		// Above average TCP:
		// size = 2000

		// Optimal Benchmark:
		// size = 4000

		// Huge:
		// size = 10000

		// Biggest (64KB):
		// size = 65533

		d    = []byte(strings.Repeat("a", size))
		data = append(d, ':', '\n')
	)

	for {
		// select {
		// case <-timer.C:
		// fmt.Println("stopping!")
		// runtime.GC()
		// return

		// default:
		// }

		// Write the delimited messages to the buffer
		// line, err = conn.Wr9ite(data)
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
