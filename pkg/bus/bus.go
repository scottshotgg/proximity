package bus

import "github.com/scottshotgg/proximity/pkg/listener"

type (
	// Bus ...
	Bus interface {
		Insert(msg *listener.Msg) error
		Remove() (*listener.Msg, error)

		Close() error
	}
)
