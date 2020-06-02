package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/scottshotgg/proximity/pkg/bus/channel"
	reciever "github.com/scottshotgg/proximity/pkg/receiver"
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

		c          = channel.New(100)
		r          = reciever.New(c)
		route      = "ur_mom"
		lisID, err = r.Attach(route)
	)

	if err != nil {
		log.Fatalln("err attaching", err)
	}

	fmt.Println("listenerID:", lisID)

	for i := 0; i < 1000; i++ {
		var (
			contents  = strconv.Itoa(i)
			blob, err = json.Marshal(&reciever.Msg{
				Route:    route,
				Contents: contents,
			})
		)

		if err != nil {
			log.Fatalln("err marshaling:", err)
		}

		err = c.Insert(string(blob))
		if err != nil {
			log.Fatalln("err:", err)
		}

		fmt.Println("inserted into:", string(blob))
		// // }

		// msg, err = c.Remove()
		// if err != nil {
		// 	log.Fatalln("err:", err)

		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("done")

	// log.Println("msg:", msg)
}
