package node

// TODO: change Node to be an interface

type (
	Node interface {
		Send() error
		Attach() error
	}
)
