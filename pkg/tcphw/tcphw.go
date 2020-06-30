package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/scottshotgg/proximity/pkg/tcphw/client"
	"github.com/scottshotgg/proximity/pkg/tcphw/server"
)

func main() {
	var serverFlag = flag.Bool("server", false, "")
	var addrFlag = flag.String("addr", "localhost:9090", "")
	var timesFlag = flag.Int("times", 1, "")

	flag.Parse()

	fmt.Println("serverFlag", *serverFlag)
	fmt.Println("addrFlag", *addrFlag)

	var wg = &sync.WaitGroup{}
	wg.Add(*timesFlag)

	for i := 0; i < *timesFlag; i++ {
		if *serverFlag {
			go server.Start()
		} else {
			go client.Start(*addrFlag)
		}
	}

	wg.Wait()
}
