package main

import (
	"flag"
	"fmt"

	"github.com/scottshotgg/proximity/pkg/tcphw/client"
	"github.com/scottshotgg/proximity/pkg/tcphw/server"
)

func main() {
	var serverFlag = flag.Bool("server", false, "")
	var addrFlag = flag.String("addr", "localhost:9090", "")

	flag.Parse()

	fmt.Println("serverFlag", *serverFlag)
	fmt.Println("addrFlag", *addrFlag)

	if *serverFlag {
		server.Start()
	} else {
		client.Start(*addrFlag)
	}
}
