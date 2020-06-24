package sender

import "github.com/scottshotgg/proximity/pkg/listener"

type (
	// Sender ...
	Sender interface {
		Open() error
		Close() error

		// Discover(nodes []string) ([]string, error)

		Send(msg *listener.Msg) error
	}
)
