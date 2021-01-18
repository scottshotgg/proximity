package client

import (
	"log"
	"sync"
	"time"

	"github.com/inhies/go-bytesize"

	"github.com/go-redis/redis"
)

const (
	B  = 1
	KB = 1024 * B

	// amount = 131072
	amount = 512 * KB
)

type statsMeta struct {
	countBytes int64
	count      int64
	total      int64
	totalBytes int64
}

type tcpClient struct {
	r    *redis.Client
	meta statsMeta
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

	// 	fmt.meta.Println("Stats:")
	// 	fmt.meta.Println("Avg:", bytesize.New(avg))

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
			log.Printf(name+" Count: %v\n", t.meta.count)
			log.Printf(name+" Bytes: %v\n", bytesize.New(float64(t.meta.countBytes)))
			t.meta.totalBytes += t.meta.countBytes
			t.meta.total += t.meta.count
			t.meta.countBytes = 0
			t.meta.count = 0
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
