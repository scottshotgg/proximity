package reciever

import (
	"github.com/scottshotgg/proximity/pkg/listener"
)

type (
	// Reciever ...
	Reciever interface {
		Open() error
		Close() error

		Attach(lis listener.Listener) error
	}
)
