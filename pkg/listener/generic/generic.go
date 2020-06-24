package generic

import (
	"errors"

	"github.com/scottshotgg/proximity/pkg/listener"
)

type (
	handler = func(msg *listener.Msg) error

	Generic struct {
		id     string
		route  string
		handle handler
		closed bool
	}
)

func New(id string, route string, handle handler) (listener.Listener, error) {
	if handle == nil {
		return nil, errors.New("handle cannot be nil")
	}

	return &Generic{
		id:     id,
		route:  route,
		handle: handle,
	}, nil
}

func (g *Generic) ID() string {
	return g.id
}

func (g *Generic) Route() string {
	return g.route
}

func (g *Generic) Handle(msg *listener.Msg) error {
	if g.closed {
		return errors.New("handle closed")
	}

	return g.handle(msg)
}

func (g *Generic) Close() error {
	g.closed = true

	return nil
}
