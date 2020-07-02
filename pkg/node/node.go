package node

type Msg struct {
	Route    string
	Contents []byte
}

type Node interface {
	// Close() error

	Send(m *Msg)
	Stream() chan<- *Msg
	// Subscribe(route string) (chan *listener.Msg, string, error)
	Join(route string) <-chan *Msg
}
