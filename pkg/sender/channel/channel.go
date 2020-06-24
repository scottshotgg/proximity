package channel

import (
	"context"
	"errors"
	"sync"

	"github.com/scottshotgg/proximity/pkg/bus"
	"github.com/scottshotgg/proximity/pkg/listener"
	"github.com/scottshotgg/proximity/pkg/sender"
)

type (
	Source struct {
		ctx    context.Context
		closed bool
		mut    *sync.RWMutex
		b      bus.Bus
	}
)

func New(ctx context.Context, b bus.Bus) sender.Sender {
	return &Source{
		ctx: ctx,
		mut: &sync.RWMutex{},
		b:   b,
	}
}

// Open ...
func (s *Source) Open() error {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.closed = false
	// TODO: should this also start send?

	return nil
}

// Close ...
func (s *Source) Close() error {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.closed = true
	// TODO: should this also stop send?

	return nil
}

func (s *Source) Send(msg *listener.Msg) error {
	// var err error

	select {
	case <-s.ctx.Done():
		// err = s.ctx.Err()
		return errors.New("sender closed")

	default:
		if !s.closed {
			go s.b.Insert(msg)
			// err =
		}
	}

	return nil
}
