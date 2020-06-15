package grpc

import (

	// grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	// grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"

	// grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	"github.com/scottshotgg/proximity/pkg/bus"
	channel_bus "github.com/scottshotgg/proximity/pkg/bus/channel"
	"github.com/scottshotgg/proximity/pkg/listener"
	"github.com/scottshotgg/proximity/pkg/recv"
	channel_recv "github.com/scottshotgg/proximity/pkg/recv/channel"
	"github.com/scottshotgg/proximity/pkg/sender"
	channel_sender "github.com/scottshotgg/proximity/pkg/sender/channel"
	"google.golang.org/grpc"
)

type (
	Node struct {
		b          bus.Bus
		s          sender.Sender
		r          recv.Recv
		listeners  map[string]listener.Listener
		grpcServer *grpc.Server
		nodes      map[string]struct{}
	}
)

/*
	Send(Node_SendServer) error
	Attach(*AttachReq, Node_AttachServer) error
*/

// TODO: make something here that takes the individual components
func New() *Node {
	var b = channel_bus.New(100)

	return &Node{
		listeners: map[string]listener.Listener{},
		nodes:     map[string]struct{}{},

		b: b,
		s: channel_sender.New(b),
		r: channel_recv.New(b),
	}
}
