package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/inhies/go-bytesize"
)

var countBytes int64
var count int64
var total int64
var totalBytes int64

func Start() {
	addr, err := net.ResolveTCPAddr("tcp", ":9090")
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	conn, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln("err ListenTCP:", err)
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
			fmt.Printf("Count: %v\n", count)
			fmt.Printf("Bytes: %v\n", bytesize.New(float64(countBytes)))
			totalBytes += countBytes
			total += count
			countBytes = 0
			count = 0
		}
	}()

	var j int

	for {
		c, err := conn.AcceptTCP()
		if err != nil {
			log.Fatalln("err AcceptTCP:", err)
		}

		go handle(j, c)
		j++
	}
}

func handle(id int, c *net.TCPConn) {
	// c.SetNoDelay(true)
	// c.SetKeepAlive(true)
	c.SetReadBuffer(64 * 1024)
	// c.SetReadDeadline(time.Now().Add(30000000 * time.Second))

	// var count int64
	// var ticker = time.NewTicker(1 * time.Second)

	// go func() {
	// 	for range ticker.C {
	// 		var i = count
	// 		count = 0
	// 		total += i

	// 		fmt.Printf("Count %d: %v\n", id, bytesize.New(float64(i)))
	// 		// fmt.Println("Total,", total)
	// 	}
	// }()

	// var br = bufio.NewReader(c)

	// var size = 1024 * 1024
	var buf = make([]byte, 64*1024)
	var line int
	var err error

	for {
		// br.ReadBytes('\n')
		// buf := make([]byte, size)
		// fmt.Println("Size:", br.Buffered())
		// fmt.Println("size, length:", br.Size(), br.Buffered())

		// if br.Size() > 10000 {
		line, err = c.Read(buf)
		// line, err = io.ReadFull(c, buf)
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Fatalln("err ReadString:", err)
		}

		atomic.AddInt64(&countBytes, int64(line))
		atomic.AddInt64(&count, 1)
		// }
	}
}
