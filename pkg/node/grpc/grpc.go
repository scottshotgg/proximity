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
	var req, err = srv.Recv()
	if err != nil {
		return err
	}

	// var errChan = make(chan error)

	// go func() {
	for {
		// 	for _, route := range req.GetRoutes() {
		pub, err := g.n.Publish(req.GetRoutes()[0])
		if err != nil {
			// errChan <- err
			// return
			log.Fatalln("err g.n.Publish:", err)
		}

		pub <- req.GetContents()
		// }

		req, err = srv.Recv()
		if err != nil {
			if err == io.EOF {
				continue
			}

			// errChan <- err
			// return
			log.Fatalln("err srv.Recv:", err)
		}
	}
	// }()

	// return <-errChan
	return nil
}

func (g *grpcNode) Subscribe(req *buffs.SubscribeReq, srv buffs.Node_SubscribeServer) error {
	var topics = req.GetTopics()

	// Subscribe to the topics
	// TODO: fix multi subscription later; for now just use the first one
	var sub, id, err = g.n.Subscribe(topics[0])
	if err != nil {
		return err
	}

	err = srv.SendHeader(metadata.New(map[string]string{
		"id": id,
	}))

	if err != nil {
		return err
	}

	// var errChan = make(chan error)

	// go func() {
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
					continue
				}

				// errChan <- err
				// return

				log.Fatalln("err srv.Send:", err)
			}
		}
	}
	// }()

	// return <-errChan
	return nil
}
