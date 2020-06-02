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
		q      chan string

		mut  *sync.RWMutex
		once sync.Once
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
func (c *Channel) Insert(msg string) error {
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
func (c *Channel) Remove() (string, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.closed {
		return "", errClosed
	}

	select {
	case msg := <-c.q:
		return msg, nil

	default:
		// Still need to analyze when this edge case could ever be hit
		return "", errRecieve
	}
}

// New ...
func New(size int) bus.Bus {
	return &Channel{
		q:   make(chan string, size),
		mut: &sync.RWMutex{},
	}
}
