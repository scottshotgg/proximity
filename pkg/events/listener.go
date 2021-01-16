package events

import (
	"errors"

	"github.com/scottshotgg/proximity/pkg/node"
)

type handler func(msg *node.Msg) error

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

func (l *Listener) Handle(msg *node.Msg) error {
	return l.handle(msg)
}
