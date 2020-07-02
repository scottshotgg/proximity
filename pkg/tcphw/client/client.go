package client

import (
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

	conn.SetWriteBuffer(amount)

	// Subtracting 2 for the newline; 64 chars: aaa...aaa\n
	// var a = strings.Repeat("a", 62) + "\n"
	// var b = strings.Repeat("b", 62) + "\n"
	// var s = "something_here\n"

	// Repeat that a thousand times: aaa...aaa\n * 1000
	// var dataa = strings.Repeat(a, KB)
	// var datas = strings.Repeat(s, 4000)
	// var _ = datas
	var line int

	// go func() {
	// 	time.Sleep(100 * time.Millisecond)

	// 	for i := 0; i < 10; i++ {
	// 		// for {
	// 		line, err = fmt.Fprint(conn, datas)
	// 		if err != nil {
	// 			log.Fatalln("err fmt.Fprintln:", err)
	// 		}

	// 		if line != len(datas) {
	// 			log.Fatalln("wtf", line, len(datas))
	// 		}

	// 		atomic.AddInt64(&countBytes, int64(line))
	// 		atomic.AddInt64(&count, 1)

	// 		time.Sleep(10 * time.Millisecond)
	// 	}
	// }()

	var size = 64 * KB
	// var size = rand.Intn(54000) + 10000
	// fmt.Println("size", size)
	var data = strings.Repeat("a", size)
	// for i := 0; i < 100; i++ {
	for {

		// var d = fmt.Sprintf()
		// // fmt.Println(d)

		line, err = fmt.Fprint(conn, data)
		if err != nil {
			log.Fatalln("err fmt.Fprintf:", err)
		}

		if line != len(data) {
			log.Fatalln("wtf", line, len(data))
		}

		atomic.AddInt64(&countBytes, int64(line))
		atomic.AddInt64(&count, 1)

		// time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)

	conn.Close()
}
