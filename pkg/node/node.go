package node

import "github.com/scottshotgg/proximity/pkg/listener"

type Node interface {
	Publish(route string) (chan<- []byte, error)
	Subscribe(route string) (<-chan *listener.Msg, error)
}
