package events

import (
	"errors"
	"sync"
)

type Eventer struct {
	input     chan []*Msg
	listeners map[string][]*Listener
	lock      *sync.RWMutex
	id        string
	route     string
	handle    handler
}

func New() *Eventer {
	var e = &Eventer{
		input:     make(chan []*Msg, 1000),
		listeners: map[string][]*Listener{},
		lock:      &sync.RWMutex{},
		// handle:    h,
	}

	e.workers(1)

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

func (e *Eventer) SendMulti(m []*Msg) error {
	e.input <- m

	return nil
}

func (e *Eventer) Send(m *Msg) error {
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

	return e.SendMulti([]*Msg{m})
}

func (e *Eventer) Listen(id, route string) chan *Msg {
	var (
		ch = make(chan *Msg, 100)

		h = func(m *Msg) error {
			// go func() {
			// fmt.Println("sending?")
			ch <- m
			// }()

			return nil
		}

		l = Listener{
			id:     id,
			route:  route,
			handle: h,
		}
	)

	e.lock.Lock()
	e.listeners[route] = append(e.listeners[route], &l)
	e.lock.Unlock()

	return ch
}
