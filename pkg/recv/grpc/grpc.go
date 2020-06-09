package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	buffs "github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/listener"
	generic "github.com/scottshotgg/proximity/pkg/listener/generic"
	"github.com/scottshotgg/proximity/pkg/recv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type (
	Server struct {
		ch recv.Recv
	}
)

func New(port int, ch recv.Recv) error {
	var (
		s = Server{
			ch: ch,
		}

		l, err = net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	)

	if err != nil {
		return err
	}

	var grpcServer = grpc.NewServer()

	buffs.RegisterRecvServer(grpcServer, &s)
	reflection.Register(grpcServer)

	return grpcServer.Serve(l)
}

var (
	errNotImplemented = status.Error(codes.Unimplemented, "not implemented")
)

func (s *Server) Open(ctx context.Context, req *buffs.OpenReq) (*buffs.OpenRes, error) {
	return nil, errNotImplemented
}

func (s *Server) Close(ctx context.Context, req *buffs.CloseReq) (*buffs.CloseRes, error) {
	return nil, errNotImplemented
}

func (s *Server) Attach(req *buffs.AttachReq, srv buffs.Recv_AttachServer) error {
	fmt.Printf("attach request: %+v\n", req)

	var msgChan = make(chan *buffs.Message, 10000)

	var lis, err = generic.New(req.GetId(), req.GetRoute(), func(msg *listener.Msg) error {
		// return srv.Send(&buffs.AttachRes{
		// 	Message: &buffs.Message{
		// 		Route:    msg.Route,
		// 		Contents: msg.Contents,
		// 	},
		// })

		msgChan <- &buffs.Message{
			Route:    msg.Route,
			Contents: msg.Contents,
		}

		return nil
	})

	if err != nil {
		return err
	}

	err = s.ch.Attach(lis)
	if err != nil {
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
			go func() {
				var err = srv.Send(&buffs.AttachRes{
					Message: msg,
				})

				if err != nil {
					log.Fatalln("err sending back:", err)
				}
			}()

		default:
		}
	}

	return nil
}
