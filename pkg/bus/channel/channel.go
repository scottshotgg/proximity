package channel

import (
	"context"
	"errors"
	"sync"

	"github.com/scottshotgg/proximity/pkg/bus"
	"github.com/scottshotgg/proximity/pkg/listener"
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
		q      chan *listener.Msg
		mut    *sync.RWMutex
		once   sync.Once
		ctx    context.Context

		// topics map[string]chan *listener.Msg
	}
)

// Close ...
func (c *Channel) Close() error {
	c.mut.Lock()

	if c.closed {
		c.mut.Unlock()
		return errClosed
	}

	c.mut.Unlock()

	c.once.Do(func() {
		close(c.q)
		c.closed = true
		// log.Fatalln("close is not implemented")
	})

	return nil
}

// Insert ...
func (c *Channel) Insert(msg *listener.Msg) error {
	// c.mut.Lock()

	// if c.closed {
	// 	// c.mut.Unlock()
	// 	return errClosed
	// }

	// c.mut.Unlock()

	//
	// var ch chan *listener.Msg
	// var ok bool

	// c.mut.RLock()
	// ch, ok = c.topics[msg.Route]
	// c.mut.RUnlock()

	// if !ok {
	// 	c.mut.Lock()
	// 	ch = make(chan *listener.Msg, 100000)
	// 	c.topics[msg.Route] = ch
	// 	c.mut.Unlock()
	// }

	// ch <- msg
	//

	if c.closed {
		return nil
	}

	select {
	case <-c.ctx.Done():

	case c.q <- msg:
	}

	// select {
	// case c.q <- msg:

	// 	// default:
	// 	// 	return errChannelFull
	// }

	return nil
}

// Remove ...
func (c *Channel) Remove() (*listener.Msg, error) {
	// c.mut.RLock()

	// if c.closed {
	// 	c.mut.RUnlock()
	// 	return nil, errClosed
	// }

	// c.mut.RUnlock()

	select {
	case <-c.ctx.Done():
		return nil, errors.New("bus closed")

	case msg := <-c.q:
		return msg, nil
	}

	// return <-c.q, nil

	// select {
	// case msg := <-c.q:
	// 	return msg, nil

	// 	// default:
	// 	// 	// Still need to analyze when this edge case could ever be hit
	// 	// 	return nil, errRecieve
	// }
}

// New ...
func New(ctx context.Context, size int) bus.Bus {
	return &Channel{
		ctx: ctx,
		q:   make(chan *listener.Msg, 1000),
		// topics: map[string]chan *listener.Msg{},
		mut: &sync.RWMutex{},
	}
}
