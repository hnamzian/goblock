package node

import (
	"context"
	"fmt"

	"github.com/hnamzian/goblock/internal/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string
	proto.UnimplementedNodeServer
}

func New() *Node {
	return &Node{
		version: "0.0.1",
	}
}

func (n *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Printf("Received version from %s\n", peer.Addr)
	return &proto.Version{Version: n.version}, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Printf("Received tx from %s\n", peer.Addr)
	return &proto.Ack{}, nil
}
