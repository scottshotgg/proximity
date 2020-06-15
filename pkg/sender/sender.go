package sender

type (
	// Sender ...
	Sender interface {
		Open() error
		Close() error

		// Discover(nodes []string) ([]string, error)

		Send(msg []byte) error
	}
)
