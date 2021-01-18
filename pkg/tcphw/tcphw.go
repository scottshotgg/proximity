package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/scottshotgg/proximity/pkg/tcphw/client"
	"github.com/scottshotgg/proximity/pkg/tcphw/server"
)

func main() {
	var serverFlag = flag.Bool("server", false, "")
	var addrFlag = flag.String("addrs", "localhost:9090", "")
	var senderFlag = flag.Bool("sender", false, "")
	var routeFlag = flag.String("route", "*", "")
	var timesFlag = flag.Int("times", 1, "")

	flag.Parse()

	fmt.Println("Server:", *serverFlag)
	fmt.Println("Serving on:", *addrFlag)

	var addrSplit = strings.Split(*addrFlag, ",")

	var n = server.New()

	time.Sleep(1 * time.Second)

	var wg = &sync.WaitGroup{}
	wg.Add(*timesFlag)

	for _, addr := range addrSplit {
		if *serverFlag == true {
			go n.Start(addr)
		} else {
			go func() {
				// TODO: uri here
				var node, err = client.New()
				if err != nil {
					log.Fatalln("err:", err)
				}

				node.Start(addr, *timesFlag, *senderFlag, *routeFlag)
			}()
		}
	}

	wg.Wait()
}
