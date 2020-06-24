package channel

import (
	"sync"

	"github.com/scottshotgg/proximity/pkg/bus"
	"github.com/scottshotgg/proximity/pkg/listener"
	"github.com/scottshotgg/proximity/pkg/sender"
)

type (
	Source struct {
		closed bool
		mut    *sync.RWMutex
		b      bus.Bus
	}
)

func New(b bus.Bus) sender.Sender {
	return &Source{
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
	return s.b.Insert(msg)
}
