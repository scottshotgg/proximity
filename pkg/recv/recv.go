package recv

import (
	"github.com/scottshotgg/proximity/pkg/listener"
)

type (
	// Recv ...
	Recv interface {
		Open() error
		Close() error

		Attach(lis listener.Listener) error
	}
)
