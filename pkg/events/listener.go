package events

import (
	"errors"
)

type handler func(msg *Msg) error

type Listener struct {
	id     string
	route  string
	handle handler
}

func NewListener(id string, route string, handle handler) (*Listener, error) {
	if handle == nil {
		return nil, errors.New("handle cannot be nil")
	}

	return &Listener{
		id:     id,
		route:  route,
		handle: handle,
	}, nil
}

func (l *Listener) Handle(msg *Msg) error {
	return l.handle(msg)
}
