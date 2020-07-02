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
	"github.com/scottshotgg/proximity/pkg/node"
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

var (
	msgChan = make(chan []*node.Msg, 1000)
)

func handle(id int, c *net.TCPConn) {
	// Set the TCP window to 64KB
	c.SetReadBuffer(amount)

	var (
		br = bufio.NewReader(c)

		// This channel amount gives a considerable increase
		// We can always make it variable at run time based on memory size
		parseChan = make(chan []byte, 100000)
	)

	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for range msgChan {
	// 		}
	// 	}()
	// }

	for i := 0; i < 10; i++ {
		go func() {
			// // - Parse the frame for messages; these are delineated with ':' for now
			// // - Send to workers
			// // - Update state and send messages

			// // for frame := range sendch {
			// // fmt.Println(len(strings.Split(string(frame), ":")))
			// // }

			// // We can definitely make some very simple improvements to
			// // drastically reduce the allocations here
			// var delin byte = ':'

			// for frame := range parseChan {
			// 	var msgs []*node.Msg

			// 	var lastIndex int
			// 	for i, b := range frame {
			// 		if b == delin {
			// 			// - Capture indicies of delineations
			// 			// - Build Msg array
			// 			msgs = append(msgs, &node.Msg{
			// 				Route:    "a",
			// 				Contents: frame[lastIndex:i],
			// 			})

			// 			lastIndex = i + 1
			// 		}
			// 	}

			// 	msgChan <- msgs
			// }

			for range parseChan {
				// var split = strings.Split(string(frame), ":")
			}
		}()
	}

	// var r, w = io.Pipe()

	// go func() {
	// 	var msgReader = bufio.NewReader(r)

	// 	for {
	// 		var b, err = msgReader.ReadBytes(':')
	// 		if err != nil {
	// 			log.Fatalln("err msgReader.ReadBytes():", err)
	// 		}

	// 		parseChan <- b
	// 	}
	// }()

	// var brw = bufio.NewWriter(w)

	for {
		// Read in a 'frame' of messages; these are delineated by newlines
		b, err := br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Fatalln("err ReadString:", err)
		}

		// brw.Write(b)

		// Send to parsers
		parseChan <- b

		atomic.AddInt64(&countBytes, int64(len(b)))
		atomic.AddInt64(&count, 1)
	}
}
