package local

import (
	"errors"
	"sync"
	"time"

	"github.com/scottshotgg/proximity/pkg/node"
)

type Local struct {
	// b         bus.Bus
	// r         recv.Recv
	listeners map[string][]chan *node.Msg
	// nodes     map[string]struct{}
	// ctx       context.Context
	// cancel    context.CancelFunc
	// closed    bool

	// bus chan []byte

	// routes map[string]map[string]chan []byte

	joiners chan *Joiner
	input   chan *node.Msg

	lock *sync.RWMutex
}

type Joiner struct {
	Route  string
	Output chan *node.Msg
}

var (
	ErrIDTaken = errors.New("ID already taken")
)

const chanSize = 1000

func New() *Local {
	// var ctx, cancel = context.WithCancel(context.Background())

	// var b = channel_bus.New(ctx, 1000)

	var l = &Local{
		listeners: map[string][]chan *node.Msg{},
		// nodes:     map[string]struct{}{},

		// b:      b,
		// ctx:    ctx,
		// cancel: cancel,
		// r:      channel_recv.New(ctx, b),
		lock:  &sync.RWMutex{},
		input: make(chan *node.Msg, chanSize),

		joiners: make(chan *Joiner, chanSize),
	}

	go l.eventLoop()

	return l
}

func (l *Local) Join(route string) <-chan *node.Msg {
	var ch = make(chan *node.Msg, chanSize)

	l.joiners <- &Joiner{
		Route:  route,
		Output: ch,
	}

	return ch
}

func (l *Local) Send(m *node.Msg) {
	l.input <- m
}

func (l *Local) Stream() chan<- *node.Msg {
	return l.input
}

func (l *Local) eventLoop() {
	go func() {
		for {
			select {
			// Check if anyone wants to join
			case j := <-l.joiners:
				l.lock.Lock()
				l.listeners[j.Route] = append(l.listeners[j.Route], j.Output)
				l.lock.Unlock()
			}

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			// Process the incoming messages
			case msg := <-l.input:
				l.lock.RLock()
				var listeners = l.listeners[msg.Route]
				l.lock.RUnlock()

				for _, listener := range listeners {
					// go func(l chan *Msg) {
					listener <- msg
					// }(listener)
				}
			}
		}
	}()
}
