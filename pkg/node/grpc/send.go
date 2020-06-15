package grpc

import (
	"encoding/json"

	"github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/listener"
)

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
