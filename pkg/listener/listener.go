package listener

// TODO: make the listener take a Recv or a Node

type (
	Listener interface {
		ID() string
		Route() string
		Handle(msg *Msg) error
	}

	Msg struct {
		// TODO: Think of something better later
		Route    string
		Contents []byte
	}
)
