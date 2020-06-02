package sender

type (
	// Sender ...
	Sender interface {
		Open() error
		Close() error

		Send(msg string) error
	}
)
