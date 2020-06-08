package main

import (
	"fmt"
	"log"

	bus "github.com/scottshotgg/proximity/pkg/bus/channel"
	"github.com/scottshotgg/proximity/pkg/listener/echo"
	recv "github.com/scottshotgg/proximity/pkg/receiver/channel"
	channel_sender "github.com/scottshotgg/proximity/pkg/sender/channel"
	sender "github.com/scottshotgg/proximity/pkg/sender/grpc"
)

// func main() {
// 	var (
// 		err error
// 		msg string

// 		c = channel.New(100)
// 	)

// 	// for i := 0; i < 1000; i++ {
// 	err = c.Insert("something here")
// 	if err != nil {
// 		log.Fatalln("err:", err)
// 	}

// 	fmt.Println("inserted")
// 	// }

// 	msg, err = c.Remove()
// 	if err != nil {
// 		log.Fatalln("err:", err)
// 	}

// 	log.Println("msg:", msg)
// }

func main() {
	var (
		// msg string
		// err error

		b       = bus.New(100)
		r       = recv.New(b)
		route1  = "ur_mom"
		route2  = "ur_dad"
		l1, err = echo.New(route1)
	)

	if err != nil {
		log.Fatalln("err creating new echo listener", err)
	}

	err = r.Attach(l1)
	if err != nil {
		log.Fatalln("err attaching", err)
	}

	l2, err := echo.New(route2)
	if err != nil {
		log.Fatalln("err attaching", err)
	}

	err = r.Attach(l2)
	if err != nil {
		log.Fatalln("err attaching", err)
	}

	fmt.Println("listener1 ID:", l1.ID())
	fmt.Println("listener2 ID:", l2.ID())

	err = sender.New(5001, channel_sender.New(b))
	if err != nil {
		log.Fatalln("err created sender:", err)
	}

	// for i := 0; i < 1000; i++ {
	// 	var (
	// 		contents = strconv.Itoa(i)
	// 		msg      = listener.Msg{
	// 			Route:    route1,
	// 			Contents: contents,
	// 		}
	// 	)

	// 	if i%2 == 0 {
	// 		msg.Route = route2
	// 	}

	// 	if i%5 == 0 {
	// 		msg.Route = recv.RouteAll
	// 	}

	// 	err = s.Send(&msg)
	// 	if err != nil {
	// 		log.Fatalln("err:", err)
	// 	}

	// 	fmt.Printf("inserted into: %+v\n", msg)
	// 	// // }

	// 	// msg, err = c.Remove()
	// 	// if err != nil {
	// 	// 	log.Fatalln("err:", err)

	// 	time.Sleep(200 * time.Millisecond)
	// }

	// fmt.Println("done")

	// // log.Println("msg:", msg)
}
