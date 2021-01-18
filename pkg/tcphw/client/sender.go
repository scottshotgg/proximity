package client

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"runtime"
	"sync/atomic"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

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

		bw = bufio.NewWriter(conn)
	)

	// Sender
	err = bw.WriteByte('1')
	if err != nil {
		log.Fatalln("br.WriteByte():", err)
	}

	err = bw.Flush()
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

	data = bytes.Repeat(data, 128)

	// var todo = Todo{
	// 	ID:   uuid.New().String(),
	// 	Text: "blah blah",
	// 	Done: true,
	// }

	// input, err := reader.ReadBytes('\n')
	// if err != nil {
	// 	log.Fatalln("err reader.ReadBytes():", err)
	// }

	for {
		// todo.ID = uuid.New().String()

		// blob, err := json.Marshal(todo)
		// if err != nil {
		// 	log.Fatalln("err marshal:", err)
		// }

		// var msg = "a:" + string(blob) + "\n"

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
		line, err = bw.WriteString(string(data))
		if err != nil {
			log.Println("err bw.Write:", err)
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
		// if line != len(msg) {
		// 	log.Fatalln("wtf dude", line, len(msg))
		// }

		atomic.AddInt64(&t.meta.countBytes, int64(line))
		atomic.AddInt64(&t.meta.count, 1)
	}
}
