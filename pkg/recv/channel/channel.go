package reciever

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/scottshotgg/proximity/pkg/bus"
	"github.com/scottshotgg/proximity/pkg/listener"
	"github.com/scottshotgg/proximity/pkg/recv"
)

const (
	// RouteAll ...
	RouteAll = "*"

	// RouteNoOp ...
	RouteNoOp = ""
)

var (
	errClosed = errors.New("reciever is closed")
)

type (
	// Sink ...
	Sink struct {
		closed    bool
		mut       *sync.RWMutex
		listeners map[string][]listener.Listener
		b         bus.Bus
	}
)

// func NewMsg(route, contents string) *Msg {
// 	return &Msg{
// 		Route:    route,
// 		Contents: contents,
// 	}
// }

// New ...
func New(b bus.Bus) recv.Recv {
	var s = Sink{
		mut:       &sync.RWMutex{},
		listeners: map[string][]listener.Listener{},
		b:         b,
	}

	go func(s *Sink) {
		var err = s.recv()
		if err != nil {
			log.Fatalln("err recv:", err)
		}
	}(&s)

	return &s
}

// Open ...
func (s *Sink) Open() error {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.closed = false
	// TODO: should this also start recv?

	return nil
}

// Close ...
func (s *Sink) Close() error {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.closed = true
	// TODO: should this also stop recv?

	return nil
}

// TODO: this needs to be a background function
// Recv ...
func (s *Sink) recv() error {
	var msgChan = make(chan []byte, 2000000)

	go func() {
		for {
			// time.Sleep(500 * time.Millisecond)

			var m, err = s.b.Remove()
			if err != nil {
				// log.Fatalln("err removing:", err)
				continue
			}

			msgChan <- m
		}
	}()

	// TODO: change this to use a future and return a channel on Recieve
	// have Sync and Async
	var workChan = make(chan struct{}, 10)

	for {
		s.mut.Lock()

		if s.closed {
			return errClosed
		}

		s.mut.Unlock()

		// Reserve a spot
		workChan <- struct{}{}

		go func() {
			// Relinquish the spot
			defer func() {
				_ = <-workChan
			}()

			// Put all below into another worker func or something

			var msg listener.Msg

			select {
			case m := <-msgChan:
				var err = json.Unmarshal([]byte(m), &msg)
				if err != nil {
					log.Fatalln("err unmarshal", err)
					// TODO: handle
					return
				}

			// TODO: timeout, this may be better as a context
			case <-time.After(1 * time.Second):
				// log.Println("err timeout")
				return
			}

			switch msg.Route {
			case RouteAll:
				go func(m *listener.Msg) {
					log.Println("broadcast")
					var err = s.broadcast(m)
					if err != nil {
						log.Println("err broadbasting", err)
					}
				}(&msg)

			case RouteNoOp:
				log.Println("noop")
				return

			default:
				var listeners, ok = s.listeners[msg.Route]
				if !ok {
					log.Println("could not find listener, tossing")
					// TODO: implement default behavior
					return
				}

				for _, l := range listeners {
					go func(l listener.Listener) {
						var err = l.Handle(&msg)
						if err != nil {
							// TODO: do something here... maybe internally queue?
						}
					}(l)
				}
			}
		}()
	}
}

func (s *Sink) broadcast(msg *listener.Msg) error {
	// Range over every route ...
	for _, route := range s.listeners {
		go func(route []listener.Listener) {
			// For each listener on that route ...
			for _, lis := range route {
				go func(l listener.Listener) {
					var err = l.Handle(msg)
					if err != nil {
						log.Fatalln("err handling", err)
					}
				}(lis)
			}
		}(route)
	}

	// return errors.New("broadcast is not implemented")
	return nil
}

// Attach adds a listener to the reciever
func (s *Sink) Attach(lis listener.Listener) error {
	// TODO: need to have a mutex specifically for this

	var (
		id    = lis.ID()
		route = lis.Route()
	)

	// TODO: need to check this route
	var r, _ = s.listeners[route]
	// if !ok {
	// 	// If we didn't find it
	// 	// s.listeners[route] =
	// 	r = []*Listener{lis}
	// } else {
	// 	// If we did find it
	// 	r = append(r, lis)
	// }

	s.listeners[route] = append(r, lis)

	fmt.Printf("Attached listener, ID:%s\n", id)

	return nil
}
