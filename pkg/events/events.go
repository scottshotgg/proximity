package events

import (
	"errors"
	"sync"
	"time"

	"github.com/scottshotgg/proximity/pkg/node"
)

type Eventer struct {
	input     chan []*node.Msg
	listeners map[string][]*Listener
	lock      *sync.RWMutex
	id        string
	route     string
	handle    handler
}

func New() *Eventer {
	var e = &Eventer{
		input:     make(chan []*node.Msg, 1000),
		listeners: map[string][]*Listener{},
		lock:      &sync.RWMutex{},
		// handle:    h,
	}

	e.workers(25)

	return e
}

func (e *Eventer) workers(amount int) {
	for i := 0; i < amount; i++ {
		go func() {
			for msgs := range e.input {
				// TODO: can probably build up the messages here
				for _, msg := range msgs {
					e.lock.RLock()
					var listeners, ok = e.listeners[msg.Route]
					e.lock.RUnlock()

					if !ok {
						return
					}

					for _, l := range listeners {
						l.Handle(msg)
					}
				}
			}
		}()
	}
}

var errUnableToRoute = errors.New("unable to resolve route")

func (e *Eventer) SendMulti(m []*node.Msg) error {
	e.input <- m

	return nil
}

func (e *Eventer) Stream() chan<- *node.Msg {
	var ch = make(chan *node.Msg, 1000)

	go func() {
		var ticker = time.NewTicker(1 * time.Second)
		var msgs = []*node.Msg{}

		for {
			select {
			case <-ticker.C:
				e.input <- msgs
				msgs = []*node.Msg{}

			case msg := <-ch:
				msgs = append(msgs, msg)
			}
		}
	}()

	return ch
}

func (e *Eventer) Send(m *node.Msg) error {
	// e.lock.RLock()
	// var listeners, ok = e.listeners[m.Route]
	// e.lock.RUnlock()

	// if !ok {
	// 	// fmt.Println("not ok")
	// 	return errUnableToRoute
	// }

	// go func() {
	// 	for _, l := range listeners {
	// 		l.Handle(m)
	// 	}
	// }()

	return e.SendMulti([]*node.Msg{m})
}

func (e *Eventer) Listen(route string) <-chan *node.Msg {
	var (
		ch = make(chan *node.Msg, 100)

		h = func(m *node.Msg) error {
			// go func() {
			// fmt.Println("sending?")
			ch <- m
			// }()

			return nil
		}

		l = Listener{
			// id:     id,
			route:  route,
			handle: h,
		}
	)

	e.lock.Lock()
	e.listeners[route] = append(e.listeners[route], &l)
	e.lock.Unlock()

	return ch
}
