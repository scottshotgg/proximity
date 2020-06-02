package reciever

import (
	"fmt"
)

type (
	Listener struct {
		id    string
		route string
	}
)

// TODO: idk think about this
func (l *Listener) Send(msg string) error {
	fmt.Printf("Event!\n\tID: %s\n\tRoute: %s\n\tMessage: %s\n\n", l.id, l.route, msg)

	return nil
}
