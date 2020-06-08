package listener

type (
	Listener interface {
		ID() string
		Route() string
		Handle(msg *Msg) error
	}

	Msg struct {
		// TODO: Think of something better later
		Route    string
		Contents string
	}
)
