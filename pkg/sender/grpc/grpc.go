package grpc

import (
	"context"
	"net"
	"strconv"

	buffs "github.com/scottshotgg/proximity/pkg/buffs"
	"github.com/scottshotgg/proximity/pkg/listener"
	"github.com/scottshotgg/proximity/pkg/sender"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type (
	Server struct {
		ch sender.Sender
	}
)

func New(port int, ch sender.Sender) error {
	var (
		s = Server{
			ch: ch,
		}

		l, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	)

	if err != nil {
		return err
	}

	var grpcServer = grpc.NewServer()

	buffs.RegisterSenderServer(grpcServer, &s)
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

func (s *Server) Send(ctx context.Context, req *buffs.SendReq) (*buffs.SendRes, error) {
	var err = s.ch.Send(&listener.Msg{
		Route:    req.GetMsg().GetRoute(),
		Contents: req.GetMsg().GetContents(),
	})

	if err != nil {
		return nil, err
	}

	return &buffs.SendRes{}, nil
}
