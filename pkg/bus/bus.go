package bus

type (
	// Bus ...
	Bus interface {
		Insert(msg string) error
		Remove() (string, error)

		Close() error
	}
)
