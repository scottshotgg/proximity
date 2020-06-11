package node

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/bus"
	channel_bus "github.com/scottshotgg/proximity/pkg/bus/channel"
	"github.com/scottshotgg/proximity/pkg/listener"
	generic_listener "github.com/scottshotgg/proximity/pkg/listener/generic"
	"github.com/scottshotgg/proximity/pkg/recv"
	channel_recv "github.com/scottshotgg/proximity/pkg/recv/channel"
	"github.com/scottshotgg/proximity/pkg/sender"
	channel_sender "github.com/scottshotgg/proximity/pkg/sender/channel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Node struct {
		b          bus.Bus
		s          sender.Sender
		r          recv.Recv
		listeners  map[string]listener.Listener
		grpcServer *grpc.Server
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
		b: b,
		s: channel_sender.New(b),
		r: channel_recv.New(b),
	}
}

func (n *Node) Send(srv buffs.Node_SendServer) error {
	for {
		var req, err = srv.Recv()
		if err != nil {
			// return errors.Wrap("sender.Recv:", err)
			return err
		}

		// TODO: ideally gofunc this, maybe have a bunch of workers here and make a feedback loop for that?
		blob, err := json.Marshal(listener.Msg{
			Route:    req.GetMsg().GetRoute(),
			Contents: req.GetMsg().GetContents(),
		})
		if err != nil {
			// return errors.Wrap("proto.Marshal:", err)
			return err
		}

		err = n.s.Send(blob)
		if err != nil {
			// return errors.Wrap("n.s.Send:", err)
			return err
		}
	}
}

func (n *Node) Attach(req *buffs.AttachReq, srv buffs.Node_AttachServer) error {
	fmt.Printf("attach request: %+v\n", req)

	var lis, err = generic_listener.New(req.GetId(), req.GetRoute(), func(msg *listener.Msg) error {
		return srv.Send(&buffs.AttachRes{
			Message: &buffs.Message{
				Route:    msg.Route,
				Contents: msg.Contents,
			},
		})
	})

	if err != nil {
		// return errors.Wrap("generic_listener.New:", err)
		return err
	}

	err = n.r.Attach(lis)
	if err != nil {
		// return errors.Wrap("n.r.Attach:", err)
		return err
	}

	// err = srv.SendHeader(metadata.MD{
	// 	"id":    []string{req.Id},
	// 	"route": []string{req.Route},
	// })

	// if err != nil {
	// 	return err
	// }

	// for {
	// 	select {
	// 	case msg := <-msgChan:
	// 		go func() {
	// 			var err = srv.Send(&buffs.AttachRes{
	// 				Message: msg,
	// 			})

	// 			if err != nil {
	// 				log.Fatalln("err sending back:", err)
	// 			}
	// 		}()

	// 	default:
	// 	}
	// }
	for {
	}

	return nil
}

// TODO: reorganize this but keep it like this for now
func (n *Node) Start(port int) error {
	var l, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return nil
	}

	var grpcServer = grpc.NewServer()

	buffs.RegisterNodeServer(grpcServer, n)
	reflection.Register(grpcServer)

	return grpcServer.Serve(l)
}
