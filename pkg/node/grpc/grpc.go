package grpc

import (
	"io"

	"github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/node"
	"github.com/scottshotgg/proximity/pkg/node/local"
	"google.golang.org/grpc/metadata"
)

type grpcNode struct {
	n node.Node
}

func New() *grpcNode {
	return &grpcNode{
		n: local.New(),
	}
}

func (g *grpcNode) Publish(srv buffs.Node_PublishServer) error {
	var req, err = srv.Recv()
	if err != nil {
		return err
	}

	pub, err := g.n.Publish(req.GetRoute())
	if err != nil {
		return err
	}

	var errChan = make(chan error)

	go func() {
		for {
			pub <- req.GetContents()

			req, err = srv.Recv()
			if err != nil {
				if err == io.EOF {
					err = nil
				}

				errChan <- err
				return
			}
		}
	}()

	return <-errChan
}

func (g *grpcNode) Subscribe(req *buffs.SubscribeReq, srv buffs.Node_SubscribeServer) error {
	// Subscribe to the route
	var sub, id, err = g.n.Subscribe(req.GetRoute())
	if err != nil {
		return err
	}

	err = srv.SendHeader(metadata.New(map[string]string{
		"id": id,
	}))

	if err != nil {
		return err
	}

	var errChan = make(chan error)

	go func() {
		for {
			select {
			case msg := <-sub:
				err = srv.Send(&buffs.SubscribeRes{
					Message: &buffs.Message{
						Contents: msg.Contents,
					},
				})

				if err != nil {
					if err == io.EOF {
						err = nil
					}

					errChan <- err
					return
				}
			}
		}
	}()

	return <-errChan
}
