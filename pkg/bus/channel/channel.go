package channel

import (
	"errors"
	"sync"

	"github.com/scottshotgg/proximity/pkg/bus"
)

var (
	errClosed      = errors.New("channel is closed")
	errChannelFull = errors.New("channel is full")
	errInsert      = errors.New("could not insert to channel")

	errRecieve = errors.New("could not recieve from channel")
)

type (
	// Channel ...
	Channel struct {
		closed bool
		q      chan []byte
		mut    *sync.RWMutex
		once   sync.Once
	}
)

// Close ...
func (c *Channel) Close() error {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.closed {
		return errClosed
	}

	c.once.Do(func() {
		close(c.q)
	})

	return nil
}

// Insert ...
func (c *Channel) Insert(msg []byte) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.closed {
		return errClosed
	}

	select {
	case c.q <- msg:

		// default:
		// 	return errChannelFull
	}

	return nil
}

// Remove ...
func (c *Channel) Remove() ([]byte, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.closed {
		return nil, errClosed
	}

	select {
	case msg := <-c.q:
		return msg, nil

	default:
		// Still need to analyze when this edge case could ever be hit
		return nil, errRecieve
	}
}

// New ...
func New(size int) bus.Bus {
	return &Channel{
		q:   make(chan []byte, size),
		mut: &sync.RWMutex{},
	}
}
