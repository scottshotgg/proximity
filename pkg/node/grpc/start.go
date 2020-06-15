package grpc

import (
	"net"
	"strconv"

	"github.com/scottshotgg/proximity/pkg/buffs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: reorganize this but keep it like this for now
func (n *Node) Start(port int) error {

	// logger, _ := zap.NewProduction()

	// logger.Info("listening")

	// grpc_zap.ReplaceGrpcLoggerV2(logger)

	var grpcServer = grpc.NewServer(
	// grpc_middleware.WithUnaryServerChain(
	// 	middleware.AddClientInformation,
	// ),
	)

	buffs.RegisterNodeServer(grpcServer, n)
	reflection.Register(grpcServer)

	var l, err = net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return nil
	}

	return grpcServer.Serve(l)
}
