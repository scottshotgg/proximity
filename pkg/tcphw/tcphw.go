package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"github.com/scottshotgg/proximity/pkg/tcphw/client"
	"github.com/scottshotgg/proximity/pkg/tcphw/server"
)

func main() {
	var serverFlag = flag.Bool("server", false, "")
	var addrFlag = flag.String("addrs", "localhost:9090", "")
	var timesFlag = flag.Int("times", 1, "")

	flag.Parse()

	fmt.Println("Server:", *serverFlag)
	fmt.Println("Serving on:", *addrFlag)

	var addrSplit = strings.Split(*addrFlag, ",")

	var n = server.New()

	var wg = &sync.WaitGroup{}
	wg.Add(*timesFlag)

	for _, addr := range addrSplit {
		if *serverFlag == true {
			go n.Start(addr)
		} else {
			go client.New().Start(addr, *timesFlag, true)
		}
	}

	wg.Wait()
}
