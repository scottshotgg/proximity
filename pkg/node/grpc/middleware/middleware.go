package middleware

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// func AddClientInformation(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// myid := atomic.AddUint64(&reqID, 1)
// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, "something_here", fmt.Sprintf("%s-%06d", "blah", 777666555))
// 		h.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func AddClientInformation(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// var md, ok = metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	// md = metadata.MD{}
	// }

	peer, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatalln("could not get peer information")
	}

	fmt.Println("peer", peer)

	// fmt.Println("ctx", md)

	// os.Exit(9)
	// var newctx = peer.NewContext(ctx, &peer.Peer{
	// 	Addr: &net.TCPAddr{
	// 		IP: net.IPv4zero,
	// 	},
	// })

	// fmt.Println("newCtx", newctx)

	return handler(ctx, req)
}
