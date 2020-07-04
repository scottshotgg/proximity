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
	"github.com/scottshotgg/proximity/pkg/events"
	"github.com/scottshotgg/proximity/pkg/node"
)

var countBytes int64
var count int64
var total int64
var totalBytes int64

func Start(addr string) {
	var e = events.New()

	address, err := net.ResolveTCPAddr("tcp", addr+":9090")
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	conn, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatalln("err ListenTCP:", err)
	}

	fmt.Println("Serving on:", addr+":9090")

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

		go handle(j, e, c)
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

func handle(id int, e *events.Eventer, c *net.TCPConn) {
	// e.Listen()

	// Set the TCP window to 64KB
	c.SetReadBuffer(amount)
	c.SetWriteBuffer(amount)

	var br = bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	var b, err = br.ReadByte()
	if err != nil {
		log.Fatalln("err br.ReadByte():", err)
	}

	switch b {
	// sender
	case '1':
		sender(id, e, br)

	// recvr
	case '2':
		recver(id, e, br)
	}
}

func recver(id int, e *events.Eventer, br *bufio.ReadWriter) {

}

func sender(id int, e *events.Eventer, br *bufio.ReadWriter) {
	// This channel amount gives a considerable increase
	// We can always make it variable at run time based on memory size
	var parseChan = make(chan []byte, 10000)

	for i := 0; i < 2; i++ {
		go func() {
			for range msgChan {
			}
		}()
	}

	for i := 0; i < 2; i++ {
		go func() {
			var (
				splitter  byte = ':'
				lastIndex int
			)

			for frame := range parseChan {
				var msgs []*node.Msg

				for i, f := range frame {
					if f == splitter {
						msgs = append(msgs, &node.Msg{
							Route:    "a",
							Contents: frame[lastIndex:i],
						})

						lastIndex = i + 1
						// _ = lastIndex
					}
				}

				msgChan <- msgs

				lastIndex = 0
			}
		}()
	}

	// var r, w = io.Pipe()

	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		var msgReader = bufio.NewReader(r)

	// 		for {
	// 			var msgs []*node.Msg

	// 			for i := 0; i < 1000; i++ {
	// 				var b, err = msgReader.ReadBytes(':')
	// 				if err != nil {
	// 					log.Fatalln("err msgReader.ReadBytes():", err)
	// 				}

	// 				msgs = append(msgs, &node.Msg{
	// 					Route:    "",
	// 					Contents: b,
	// 				})
	// 			}

	// 			msgChan <- msgs
	// 		}
	// 	}()
	// }

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
