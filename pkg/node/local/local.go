package local

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/scottshotgg/proximity/pkg/bus"
	channel_bus "github.com/scottshotgg/proximity/pkg/bus/channel"
	"github.com/scottshotgg/proximity/pkg/listener"
	generic_lis "github.com/scottshotgg/proximity/pkg/listener/generic"
	"github.com/scottshotgg/proximity/pkg/node"
	"github.com/scottshotgg/proximity/pkg/recv"
	channel_recv "github.com/scottshotgg/proximity/pkg/recv/channel"
	"github.com/scottshotgg/proximity/pkg/sender"
	channel_sender "github.com/scottshotgg/proximity/pkg/sender/channel"
)

type local struct {
	b         bus.Bus
	s         sender.Sender
	r         recv.Recv
	listeners map[string]listener.Listener
	// grpcServer *grpc.Server
	nodes map[string]struct{}

	// bus chan []byte

	// routes map[string]map[string]chan []byte

	lock *sync.RWMutex
}

var (
	ErrIDTaken = errors.New("ID already taken")
)

func New() node.Node {
	var b = channel_bus.New(100)

	return &local{
		listeners: map[string]listener.Listener{},
		// nodes:     map[string]struct{}{},

		b:    b,
		s:    channel_sender.New(b),
		r:    channel_recv.New(b),
		lock: &sync.RWMutex{},
	}
}

// func (l *local) Publish(route string) (chan<- []byte, <-chan error) {
func (l *local) Publish(route string) (chan<- []byte, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	var (
		ch      = make(chan []byte, 100000000)
		errChan = make(chan error)
	)

	go func() {
		for {
			select {
			case msg := <-ch:
				var err = l.s.Send(&listener.Msg{
					Route:    route,
					Contents: msg,
				})

				if err != nil {
					errChan <- err
				}
			}
		}
	}()

	return ch, nil
}

func (l *local) Subscribe(route string) (<-chan *listener.Msg, string, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	var id, err = uuid.NewRandom()
	if err != nil {
		return nil, "", err
	}

	var ch = make(chan *listener.Msg, 100000000)

	lis, err := generic_lis.New(id.String(), route, func(msg *listener.Msg) error {
		ch <- msg

		return nil
	})

	if err != nil {
		return nil, "", err
	}

	err = l.r.Attach(lis)
	if err != nil {
		return nil, "", err
	}

	return ch, id.String(), nil
}
