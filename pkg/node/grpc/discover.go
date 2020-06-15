package grpc

import (
	"context"
	"log"

	"github.com/scottshotgg/proximity/pkg/buffs"
	"google.golang.org/grpc/peer"
)

// TODO: this seems stupid for the most simple case, but some form of cryptographic signature will give away the intricacies
func (n *Node) Discover(ctx context.Context, req *buffs.DiscoverReq) (*buffs.DiscoverRes, error) {
	// var (
	// 	nodes = req.GetNodes()
	// 	ok    bool
	// )

	// for i, node := range nodes {
	// 	_, ok = n.nodes[node]
	// 	if ok {
	// 		nodes = append(nodes[:i], nodes[i+1:]...)
	// 	}
	// }

	// fmt.Println("nodes:", nodes)

	// if len(nodes) != 0 {
	// 	log.Fatalln("we don't have these nodes")
	// }

	peer, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatalln("could not get peer information")
	}

	n.nodes[peer.Addr.String()] = struct{}{}

	return &buffs.DiscoverRes{}, nil
}
