package local

import (
	"errors"
)

type Local struct {
	// b         bus.Bus
	// r         recv.Recv
	listeners map[string][]chan *Msg
	// nodes     map[string]struct{}
	// ctx       context.Context
	// cancel    context.CancelFunc
	// closed    bool

	// bus chan []byte

	// routes map[string]map[string]chan []byte

	joiners chan *Joiner
	input   chan *Msg

	// lock *sync.RWMutex
}

type Msg struct {
	Route    string
	Contents []byte
}

type Joiner struct {
	Route  string
	Output chan *Msg
}

var (
	ErrIDTaken = errors.New("ID already taken")
)

const chanSize = 1000

func New() *Local {
	// var ctx, cancel = context.WithCancel(context.Background())

	// var b = channel_bus.New(ctx, 1000)

	var l = &Local{
		listeners: map[string][]chan *Msg{},
		// nodes:     map[string]struct{}{},

		// b:      b,
		// ctx:    ctx,
		// cancel: cancel,
		// r:      channel_recv.New(ctx, b),
		// lock:   &sync.RWMutex{},
		input: make(chan *Msg, chanSize),

		joiners: make(chan *Joiner, chanSize),
	}

	go l.eventLoop()

	return l
}

func (l *Local) Join(route string) <-chan *Msg {
	var ch = make(chan *Msg, chanSize)

	l.joiners <- &Joiner{
		Route:  route,
		Output: ch,
	}

	return ch
}

func (l *Local) Send(m *Msg) {
	l.input <- m
}

func (l *Local) Stream() chan<- *Msg {
	return l.input
}

func (l *Local) eventLoop() {
	for {
		select {
		// Check if anyone wants to join
		case j := <-l.joiners:
			l.listeners[j.Route] = append(l.listeners[j.Route], j.Output)

		// Process the incoming messages
		case msg := <-l.input:
			// var listeners, ok = l.listeners[msg.Route]
			// if !ok {
			// 	continue
			// }

			for _, listener := range l.listeners[msg.Route] {
				// go func(l chan *Msg) {
				listener <- msg
				// }(listener)
			}
		}
	}
}
