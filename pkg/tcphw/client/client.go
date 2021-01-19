package client

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/inhies/go-bytesize"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type tcpClient struct {
	r          *redis.Client
	countBytes int64
	count      int64
	total      int64
	totalBytes int64
}

func New() (*tcpClient, error) {
	var (
		uri = "redis:6379"

		// TODO: options
		client = redis.NewClient(&redis.Options{
			Addr:            uri,
			MaxRetries:      3,
			MinRetryBackoff: 1 * time.Second,
			MaxRetryBackoff: 10 * time.Second,
			// Password: "", // no password set
			// DB:       0,  // use default DB
		})

		err = client.Ping().Err()
	)

	if err != nil {
		return nil, err
	}

	return &tcpClient{
		r: client,
	}, nil
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
	// size = 510

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

	// d    = bytes.Repeat([]byte("a"), size)
	// data = append(d, ':', '\n')
	// reader = bufio.NewReader(os.Stdin)
	)

	// data = bytes.Repeat(data, 128)

	// input, err := reader.ReadBytes('\n')
	// if err != nil {
	// 	log.Fatalln("err reader.ReadBytes():", err)
	// }

	var todo = &Todo{
		Text: strings.Repeat("blah ", 10000),
		Done: true,
	}

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
		todo.ID = uuid.New().String()

		data, err := json.Marshal(todo)
		if err != nil {
			log.Fatalln("err marshal:", err)
		}

		// Write the delimited messages to the buffer
		// line, err = conn.Write(data)
		line, err = brw.Write(append(data, '\n'))
		if err != nil {
			log.Println("err brw.Write:", err)
			return
		}

		// time.Sleep(1 * time.Millisecond)

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

	// var recvChan = make(chan []byte, 10000)

	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		var ticker = time.NewTicker(1 * time.Second)

	// 		var r = t.r.Pipeline()

	// 		for {
	// 			select {
	// 			case <-ticker.C:
	// 				_, err = r.Exec()
	// 				if err != nil {
	// 					log.Fatalln("err pipeline exec:", err)
	// 				}

	// 			case d := <-recvChan:
	// 				// var msg Todo

	// 				// err = json.Unmarshal(d, &msg)
	// 				// if err != nil {
	// 				// 	log.Fatalln("err:", err, string(d))
	// 				// }

	// 				// err = r.Set(msg.ID, d, 0).Err()
	// 				// if err != nil {
	// 				// 	log.Fatalln("err set:", err)
	// 				// }

	// 			}
	// 		}
	// 	}()
	// }

	f, err := os.Create("/bin/something/here.txt")
	if err != nil {
		log.Fatalln("err:", err)
	}

	defer f.Close()

	var bufferedFile = bufio.NewWriter(f)

	_, err = io.Copy(bufferedFile, brw)
	if err != nil {
		log.Fatalln("wtf copy:", err)
	}

	// for {
	// 	runtime.Gosched()

	// 	// Write the delimited messages to the buffer
	// 	// line, err = conn.Write(data)
	// 	var data, err = brw.ReadBytes('\n')
	// 	if err != nil {
	// 		log.Fatalln("err brw.ReadBytes:", err)
	// 		return
	// 	}

	// 	// recvChan <- data

	// 	atomic.AddInt64(&t.countBytes, int64(len(data)))
	// 	atomic.AddInt64(&t.count, 1)
	// }
}
