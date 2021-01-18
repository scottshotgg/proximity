package client

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/inhies/go-bytesize"
)

type tcpClient struct {
	countBytes int64
	count      int64
	total      int64
	totalBytes int64
}

func New() *tcpClient {
	return &tcpClient{}
}

func (t *tcpClient) Start(addr string, times int, sender bool, route string) {
	// var sigChan = make(chan os.Signal)
	// signal.Notify(sigChan, os.Interrupt)

	// var start = time.Now()

	// go func() {
	// 	<-sigChan

	// 	totalBytes += countBytes
	// 	total += count

	// 	var (
	// 		elapsed = time.Now().Sub(start)
	// 		avg     = float64(totalBytes) / elapsed.Seconds()
	// 	)

	// 	fmt.Println("Stats:")
	// 	fmt.Println("Avg:", bytesize.New(avg))

	// 	os.Exit(9)
	// }()

	var ticker = time.NewTicker(1 * time.Second)

	go func() {
		var name = "Recv"
		if sender == true {
			name = "Send"
		}

		for tt := range ticker.C {
			log.Printf("t: %d", tt.Second())
			log.Printf(name+" Count: %v\n", t.count)
			log.Printf(name+" Bytes: %v\n", bytesize.New(float64(t.countBytes)))
			t.totalBytes += t.countBytes
			t.total += t.count
			t.countBytes = 0
			t.count = 0
		}
	}()

	var wg = &sync.WaitGroup{}

	var fn = t.recv

	if sender == true {
		fn = t.send
	}

	wg.Add(1)

	for i := 0; i < times; i++ {
		go fn(addr)
	}

	wg.Wait()

	ticker.Stop()
}

const (
	B  = 1
	KB = 1024 * B

	// amount = 131072
	amount = 512 * KB
)

func (t *tcpClient) send(addr string) {
	serverAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	// var timer = time.NewTimer(5 * time.Second)

	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatalln("err DialTCP:", err)
	}

	conn.SetNoDelay(true)

	// Set the TCP window to 64KB
	conn.SetWriteBuffer(amount)
	conn.SetReadBuffer(amount)

	var (
		line int

		brw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	)

	// Sender
	err = brw.WriteByte('1')
	if err != nil {
		log.Fatalln("br.WriteByte():", err)
	}

	err = brw.Flush()
	if err != nil {
		log.Fatalln("err brw.Flush():", err)
	}

	var (
		// Smallest:
		// size = 1

		// Tiny:
		// size = 5

		// Smaller:
		// size = 10

		// Small:
		// size = 100

		// Medium:
		size = 510

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
		// size = 65534

		d    = bytes.Repeat([]byte("a"), size)
		data = append(d, ':', '\n')
	// reader = bufio.NewReader(os.Stdin)
	)

	// data = bytes.Repeat(data, 128)

	// input, err := reader.ReadBytes('\n')
	// if err != nil {
	// 	log.Fatalln("err reader.ReadBytes():", err)
	// }

	for {
		runtime.Gosched()

		// select {
		// case <-timer.C:
		// fmt.Println("stopping!")
		// runtime.GC()
		// return

		// default:
		// }

		// time.Sleep(1 * time.Second)

		// Write the delimited messages to the buffer
		// line, err = conn.Write(data)
		line, err = brw.Write(data)
		if err != nil {
			log.Println("err brw.Write:", err)
			return
		}

		// err = brw.Flush()
		// if err != nil {
		// 	log.Fatalln("err brw.Flush():", err)
		// }

		// err = conn.Close()
		// if err != nil {
		// 	log.Fatalln("err conn.Close():", err)
		// }

		// // Check if the amount sent is the same as the length of the data
		// if line != len(data) {
		// 	log.Fatalln("wtf dude", line, len(data))
		// }

		atomic.AddInt64(&t.countBytes, int64(line))
		atomic.AddInt64(&t.count, 1)
	}
}

func (t *tcpClient) recv(addr string) {
	serverAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalln("err ResolveTCPAddr:", err)
	}

	// var timer = time.NewTimer(5 * time.Second)

	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatalln("err DialTCP:", err)
	}

	conn.SetNoDelay(true)

	// Set the TCP window to 64KB
	conn.SetWriteBuffer(amount)
	conn.SetReadBuffer(amount)

	var (
		brw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	)

	// Sender
	err = brw.WriteByte('2')
	if err != nil {
		log.Fatalln("br.WriteByte():", err)
	}

	err = brw.Flush()
	if err != nil {
		log.Fatalln("err brw.Flush():", err)
	}

	for {
		runtime.Gosched()

		// Write the delimited messages to the buffer
		// line, err = conn.Write(data)
		var data, err = brw.ReadBytes('\n')
		if err != nil {
			log.Println("err brw.ReadBytes:", err)
			return
		}

		// fmt.Println("data:", string(data))

		// if data[len(data)-1] != '\n' {
		// 	fmt.Println("wtf:", string(data))
		// }

		atomic.AddInt64(&t.countBytes, int64(len(data)))
		atomic.AddInt64(&t.count, 1)
	}
}
