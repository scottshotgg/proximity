package grpc

import (
	"fmt"
	"log"

	"github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/listener"
	generic_listener "github.com/scottshotgg/proximity/pkg/listener/generic"
	"google.golang.org/grpc/metadata"
)

func (n *Node) Attach(req *buffs.AttachReq, srv buffs.Node_AttachServer) error {
	fmt.Printf("attach request: %+v\n", req)

	var (
		msgChan = make(chan *buffs.Message, 1000)

		lis, err = generic_listener.New(req.GetId(), req.GetRoute(), func(msg *listener.Msg) error {
			msgChan <- &buffs.Message{
				Route:    msg.Route,
				Contents: msg.Contents,
			}

			return nil
		})
	)

	if err != nil {
		// return errors.Wrap("generic_listener.New:", err)
		return err
	}

	err = n.r.Attach(lis)
	if err != nil {
		// return errors.Wrap("n.r.Attach:", err)
		return err
	}

	err = srv.SendHeader(metadata.MD{
		"id":    []string{req.Id},
		"route": []string{req.Route},
	})

	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-msgChan:
			go func(msg *buffs.Message) {
				err = srv.Send(&buffs.AttachRes{
					Message: msg,
				})

				if err != nil {
					log.Fatalln("err sending back:", err)
				}
			}(msg)

		default:
		}
	}
}
