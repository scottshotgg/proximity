package grpc

import (
	"io"
	"log"

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
	var st = g.n.Stream()

	for {
		var req, err = srv.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			log.Fatalln("err srv.Recv:", err)

			return err
		}

		st <- &node.Msg{
			Route:    req.GetRoutes()[0],
			Contents: req.GetContents(),
		}

		// g.n.Send()
	}
}

func (g *grpcNode) Subscribe(req *buffs.SubscribeReq, srv buffs.Node_SubscribeServer) error {
	// Subscribe to the route
	var ch = g.n.Listen(req.GetTopics()[0])

	var err = srv.SendHeader(metadata.New(map[string]string{
		"id": "id",
	}))

	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-ch:
			err = srv.Send(&buffs.SubscribeRes{
				Message: &buffs.Message{
					Contents: msg.Contents,
				},
			})

			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
		}
	}
}
