package events

import (
	"errors"
	"sync"
)

type handler func(fragment string, msg *Msg) error

type Listener struct {
	id     string
	route  string
	handle handler
}

func NewListener(id string, route string, handle handler) (*Eventer, error) {
	if handle == nil {
		return nil, errors.New("handle cannot be nil")
	}

	// return &Listener{
	// 	id:     id,
	// 	route:  route,
	// 	handle: handle,
	// }, nil

	return &Eventer{
		listeners: map[string][]*Eventer{},
		lock:      &sync.RWMutex{},
	}, nil
}

func (l *Listener) Handle(fragment string, msg *Msg) error {
	return l.handle(fragment, msg)
}
