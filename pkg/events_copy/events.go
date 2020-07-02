package events

import (
	"errors"
	"fmt"
	"sync"
)

type Eventer struct {
	// input     chan *Msg
	listeners map[string][]*Eventer
	lock      *sync.RWMutex
	id        string
	route     string
	handle    handler
}

func New() *Eventer {
	var e = &Eventer{
		// input:     make(chan *Msg, 1000),
		listeners: map[string][]*Eventer{},
		lock:      &sync.RWMutex{},
		// handle:    h,
	}

	// e.workers()

	return e
}

// func (e *Eventer) workers() {
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			for m := range e.input {
// 				e.lock.RLock()
// 				var listeners, ok = e.listeners[m.Route]
// 				e.lock.RUnlock()

// 				if !ok {
// 					return
// 				}

// 				for _, l := range listeners {
// 					l.Handle(m)
// 				}
// 			}
// 		}()
// 	}
// }

var errUnableToRoute = errors.New("unable to resolve route")

func (e *Eventer) Handle(fragment string, m *Msg) error {
	// var split = strings.SplitN(fragment, "/", 2)
	// var route = split[0]
	var listeners []*Eventer
	var ok bool

	if route != "" {
		e.lock.RLock()
		listeners, ok = e.listeners[route]
		e.lock.RUnlock()

		if !ok {
			// fmt.Println("not ok")
			return errUnableToRoute
		}

		for _, l := range listeners {
			if len(split) == 1 {
				l.handle(fragment, m)
				// 		} else {
				// 			l.Handle(split[1], m)
				// 		}
				// 	}
				// } else {
				// 	for _, listeners := range e.listeners {
				// 		for _, l := range listeners {
				// 			if len(split) == 1 {
				// 				l.handle(fragment, m)
				// 			} else {
				// 				l.Handle(split[1], m)
				// 			}
			}
		}
	}

	return nil
}

// func (e *Eventer) Send(m *Msg) error {
// 	e.lock.RLock()
// 	var l, ok = e.listeners[m.Route]
// 	e.lock.RUnlock()

// 	if !ok {
// 		return errUnableToRoute
// 	}

// 	// for _, l := range listeners {
// 	l.Handle(m)
// 	// }

// 	return nil
// }

func (e *Eventer) Listen(id, route string) (*Eventer, chan *Msg) {
	var (
		ch = make(chan *Msg, 1000)

		h = func(fragment string, m *Msg) error {
			// go func() {
			// fmt.Println("sending?")
			ch <- m
			// }()

			return nil
		}

		l = Eventer{
			id:        id,
			route:     route,
			handle:    h,
			listeners: map[string][]*Eventer{},
			lock:      &sync.RWMutex{},
		}
	)

	e.lock.Lock()
	e.listeners[route] = append(e.listeners[route], &l)
	e.lock.Unlock()

	fmt.Println("ID:", id, "Network map:", e.listeners)

	return &l, ch
}
