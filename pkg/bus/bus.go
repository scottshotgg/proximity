package bus

type (
	// Bus ...
	Bus interface {
		Insert(msg []byte) error
		Remove() ([]byte, error)

		Close() error
	}
)
