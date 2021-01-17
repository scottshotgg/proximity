package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/inhies/go-bytesize"
	"github.com/scottshotgg/proximity/pkg/node"
	"github.com/scottshotgg/proximity/pkg/node/local"
)

type tcpNode struct {
	n node.Node
}

func New() *tcpNode {
	return &tcpNode{
		n: local.New(),
		// n: events.New(),
	}
}

var countBytes int64
var count int64
var total int64
var totalBytes int64

func (t *tcpNode) Start(addr string) {
	address, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	conn, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatalln("err ListenTCP:", err)
	}

	fmt.Println("Serving on:", addr)

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

		log.Println("Stats:")
		log.Println("Avg:", bytesize.New(avg))

		os.Exit(9)
	}()

	var ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			log.Printf("t: %d", t.Second())
			log.Printf("Count: %v\n", count)
			log.Printf("Bytes: %v\n", bytesize.New(float64(countBytes)))
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

		go t.handle(j, t.n, c)
		j++
	}
}

const (
	B  = 1
	KB = 1024 * B

	amount = 512 * KB
)

func (t *tcpNode) handle(id int, e node.Node, c *net.TCPConn) {
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
		// fmt.Println("sender")
		t.sender(id, e, c)

	// recvr
	case '2':
		// fmt.Println("got a recvr")
		t.recver(id, e, c)

	default:
		log.Fatalln("err b:", b)
	}
}

// This channel amount gives a considerable increase
// We can always make it variable at run time based on memory size

func (t *tcpNode) recver(id int, e node.Node, c net.Conn) {
	var brw = bufio.NewWriter(c)

	var ch = t.n.Listen("a")

	// var b = make([]byte, 64*KB)
	// var line int
	// var err error

	for msg := range ch {
		runtime.Gosched()

		// Read in a 'frame' of messages; these are delineated by newlines
		b, err := brw.Write(append(msg.Contents, '\n'))
		// var b, err = c.Write(append(msg.Contents, '\n'))
		if err != nil {
			if err == io.EOF {
				log.Println("server io.EOF")
				return
			}

			log.Fatalln("err ReadString:", err)
		}

		// brw.Flush()

		atomic.AddInt64(&countBytes, int64(b))
		// atomic.AddInt64(&countBytes, int64(len(line)))
		atomic.AddInt64(&count, 1)
	}
}

func (t *tcpNode) worker1(p <-chan []byte) {
	var (
		// splitter byte = ':'
		//
		// lastIndex int

		st = t.n.Stream()
	)

	for frame := range p {
		// for i, f := range frame {
		// 	if f == splitter {
		st <- &node.Msg{
			Route:    "a",
			Contents: frame,
		}

		// lastIndex = i + 1
		// 	}
		// }

		// lastIndex = 0
	}
}

func worker2() {
	// go func() {
	// 	var splitters = []byte{':'}

	// 	for frame := range parseChan {
	// 		var (
	// 			a    = bytes.Split(frame, splitters)
	// 			msgs []*node.Msg
	// 		)

	// 		for _, msg := range a {
	// 			msgs = append(msgs, &node.Msg{
	// 				Route:    "a",
	// 				Contents: msg,
	// 			})
	// 		}

	// 		msgChan <- msgs
	// 	}
	// }()
}

func (t *tcpNode) sender(id int, e node.Node, c net.Conn) {
	// // for i := 0; i < 2; i++ {
	// go func() {
	// 	for range msgChan {
	// 	}
	// }()
	// // }
	var p = make(chan []byte, 1000)

	defer c.Close()
	defer close(p)

	// // for i := 0; i < 2; i++ {
	go t.worker1(p)
	go t.worker1(p)
	go t.worker1(p)
	// go t.worker1()
	// go t.worker1()
	// }

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

	var brw = bufio.NewReader(c)

	// var line int
	// var err error

	for {
		var b = make([]byte, 64*KB)

		runtime.Gosched()

		// Read in a 'frame' of messages; these are delineated by newlines
		// var line, err = c.Read(b)
		var line, err = brw.Read(b)
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Println("err ReadString:", err)
			return
		}

		// brw.Write(b)

		// // Send to parsers
		p <- b

		// atomic.AddInt64(&countBytes, int64(n))
		atomic.AddInt64(&countBytes, int64(line))
		atomic.AddInt64(&count, 1)
	}
}
