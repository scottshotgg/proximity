package echo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/scottshotgg/proximity/pkg/listener"
)

type (
	// Echo implements Listener
	Echo struct {
		id    string
		route string
	}
)

// New instantiates a new Listener
func New(route string) (listener.Listener, error) {
	var id, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Echo{
		id:    id.String(),
		route: route,
	}, nil
}

func (e *Echo) ID() string {
	return e.id
}

func (e *Echo) Route() string {
	return e.route
}

func (e *Echo) Handle(msg *listener.Msg) error {
	fmt.Printf("Event! : {\n\t\"id\": \"%s\",\n\t\"route\": \"%s\"\n\t\"message\": \"%s\"\n}\n\n", e.id, e.route, msg.Contents)

	return nil
}
