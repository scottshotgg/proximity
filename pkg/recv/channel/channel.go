package reciever

import (
	"errors"
	"fmt"
	"log"
	"strings"
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

	// RouteID ...
	RouteID = "_id"

	// RouteMeta is a predefined topic that routes metadata
	RouteMeta = "_meta"
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
	var msgChan = make(chan *listener.Msg, 2000000)

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
	// var workChan = make(chan struct{}, 10)

	for {
		// Put all below into another worker func or something

		select {
		case m := <-msgChan:
			s.route(m)

		// TODO: timeout, this may be better as a context
		case <-time.After(1 * time.Second):
			// log.Println("err timeout")
			continue
		}
	}
}

func (s *Sink) route(msg *listener.Msg) {
	var split = strings.Split(msg.Route, "/")
	if len(split) < 1 {
		// We cannot do anything with this message, throw it in a log or something
		return
	}

	switch split[0] {
	case RouteAll:
		go func(m *listener.Msg) {
			log.Println("broadcast")
			var err = s.broadcast(m)
			if err != nil {
				log.Println("err broadbasting", err)
			}
		}(msg)

	case RouteNoOp:
		log.Println("noop")
		return

	case RouteID:
		fallthrough

	default:
		var (
			listeners []listener.Listener
			ok        bool
		)

		s.mut.Lock()
		listeners, ok = s.listeners[msg.Route]
		s.mut.Unlock()

		if !ok {
			log.Println("could not find listener, tossing:", msg.Route)
			// TODO: implement default behavior
			return
		}

		for _, l := range listeners {
			go func(l listener.Listener) {
				var err = l.Handle(msg)
				if err != nil {
					// TODO: do something here... maybe internally queue?
				}
			}(l)
		}
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

	var split = strings.Split(route, "/")
	if split[0] == RouteID {
		return errors.New("Cannot subscribe to ID topic")
	}

	fmt.Println("split", split, route)

	var topics = []string{
		// TODO: check whether they have ignored the ID topic
		// Listen to the ID
		RouteID + "/" + id,
	}

	if route != RouteNoOp {
		topics = append(topics, route)
	}

	// TODO: need to check this route
	s.mut.Lock()

	for _, topic := range topics {
		s.listeners[topic] = append(s.listeners[topic], lis)
	}

	s.mut.Unlock()

	fmt.Printf("Attached listener, ID:%s\n", id)
	fmt.Println("Network Map:", s.listeners)

	return nil
}
