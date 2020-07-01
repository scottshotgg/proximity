package server

import (
	"bufio"
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
			fmt.Printf("Recv Count: %v\n", count)
			fmt.Printf("Recv Bytes: %v\n", bytesize.New(float64(countBytes)))
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

const (
	B  = 1
	KB = 1024 * B

	amount = 64 * KB
)

func handle(id int, c *net.TCPConn) {
	// c.SetNoDelay(true)
	// c.SetKeepAlive(true)
	c.SetReadBuffer(amount)
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

	var br = bufio.NewReader(c)

	// var size = 1024 * 1024
	var line int
	var err error

	var sendch = make(chan []byte, 10)

	for i := 0; i < 10; i++ {
		go func() {
			for _ = range sendch {
				// // var size, err = strconv.Atoi(string(msg[0:5]))
				// // if err != nil {
				// // 	log.Fatalln("err strconv.Atoi:", err)
				// // }

				// fmt.Println("server size:", string(msg[0:5]))
				// if string(msg[0:5]) == "aaaaa" {
				// 	log.Fatalln("MESSAGE:", msg)
				// }

				// time.Sleep(1 * time.Second)
			}
		}()
	}

	for {
		// var buf = make([]byte, 5)
		// // buf := make([]byte, size)
		// // fmt.Println("Size:", br.Buffered())
		// // fmt.Println("size, length:", br.Size(), br.Buffered())

		// // if br.Size() > 10000 {
		// line, err = c.Read(buf)
		// // line, err = io.ReadFull(c, buf)
		// if err != nil {
		// 	if err == io.EOF {
		// 		return
		// 	}

		// 	log.Fatalln("err ReadString:", err)
		// }

		// fmt.Println("server buf size:", string(buf))

		// var size, err = strconv.Atoi(string(buf))
		// if err != nil {
		// 	log.Fatalln("err strconv.Atoi:", err)
		// }

		// fmt.Println("server size:", size)

		var b = make([]byte, amount)
		line, err = br.Read(b)

		// line, err = c.Read(b)
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Fatalln("err ReadString:", err)
		}

		// fmt.Println("line:", line)

		// time.Sleep(1 * time.Second)

		// for range buf {
		sendch <- b
		// }

		// strings.Split(string(buf), "\n")

		// time.Sleep(100 * time.Millisecond)

		atomic.AddInt64(&countBytes, int64(line))
		atomic.AddInt64(&count, 1)
		// }
	}
}
