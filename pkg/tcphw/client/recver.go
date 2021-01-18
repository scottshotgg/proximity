package client

import (
	"bufio"
	"log"
	"net"
	"runtime"
	"sync/atomic"
)

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

	// var recvChan = make(chan string, 10000)

	// go func() {
	for {
		runtime.Gosched()

		// var msg Todo

		// Write the delimited messages to the buffer
		// line, err = conn.Write(data)
		var data, err = brw.ReadBytes('\n')
		if err != nil {
			log.Println("err brw.ReadBytes:", err)
			return
		}

		// for data[len(data)-1] != '\n' {
		// 	var d, err = brw.ReadBytes('\n')
		// 	if err != nil {
		// 		log.Println("err brw.ReadBytes:", err)
		// 		return
		// 	}

		// 	data = append(data, d...)
		// }

		// fmt.Println("data:", data)

		// fmt.Println("len:", len(data))

		// err = json.Unmarshal(data[2:], &msg)
		// if err != nil {
		// 	fmt.Println("ERR:", err, data)
		// 	continue
		// }

		// // recvChan <- string(data[2:])

		atomic.AddInt64(&t.meta.countBytes, int64(len(data)))
		atomic.AddInt64(&t.meta.count, 1)
	}
	// }()

	// for d := range recvChan {
	// 	var dd = d[:]

	// 	var msg Todo
	// 	fmt.Println("string(d)", dd)

	// 	var err = json.Unmarshal([]byte(dd), &msg)
	// 	if err != nil {
	// 		fmt.Println("ERR:", err, dd)
	// 		continue
	// 	}

	// 	err = t.r.Set(msg.ID, dd, 0).Err()
	// 	if err != nil {
	// 		log.Fatalln("err set:", err)
	// 	}
	// }
}
