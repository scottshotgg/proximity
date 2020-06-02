package reciever

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/scottshotgg/proximity/pkg/bus"
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
	// Reciever ...
	Reciever interface {
		Open() error
		Close() error

		Attach(route string) (string, error)
	}

	// Sink ...
	Sink struct {
		closed    bool
		mut       *sync.RWMutex
		listeners map[string][]*Listener
		b         bus.Bus
	}

	Msg struct {
		// TODO: Think of something better later
		Route    string
		Contents string
	}
)

// func NewMsg(route, contents string) *Msg {
// 	return &Msg{
// 		Route:    route,
// 		Contents: contents,
// 	}
// }

// New ...
func New(b bus.Bus) Reciever {
	var s = Sink{
		mut:       &sync.RWMutex{},
		listeners: map[string][]*Listener{},
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
	var msgChan = make(chan string, 10000)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("waiting for msg ...")
			var m, err = s.b.Remove()
			if err != nil {
				fmt.Println("continue", err)
				continue
			}

			fmt.Println("got msg", m)

			msgChan <- m

			fmt.Println("passed msg")
		}
	}()

	var workChan = make(chan struct{}, 100)

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

			var msg Msg

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
				go func(m *Msg) {
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
					go func(l *Listener) {
						var err = l.Send(msg.Contents)
						if err != nil {
							// TODO: do something here... maybe internally queue?
						}
					}(l)
				}
			}
		}()
	}
}

func (s *Sink) broadcast(msg *Msg) error {
	return errors.New("broadbase is not implemented")
}

func (s *Sink) Attach(route string) (string, error) {
	// TODO: need to have a mutex specifically for this

	var id, err = uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	var idString = id.String()
	var lis = &Listener{
		id:    idString,
		route: route,
	}

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

	fmt.Printf("Attached listener, ID:%s\n", idString)

	return idString, nil
}
