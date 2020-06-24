package node

import "github.com/scottshotgg/proximity/pkg/listener"

type Node interface {
	Close() error

	Publish(route string) (chan<- []byte, error)
	Subscribe(route string) (chan *listener.Msg, string, error)
}
