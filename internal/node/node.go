package node

import (
	"context"
	"fmt"
	"sync"

	"github.com/hnamzian/goblock/internal/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	addr string

	version string

	plock sync.RWMutex
	peers map[string]proto.NodeClient

	proto.UnimplementedNodeServer
}

type Config struct {
	Addr string
}

func New(cfg *Config) *Node {
	return &Node{
		addr:    cfg.Addr,
		version: "0.0.1",
		peers:   make(map[string]proto.NodeClient),
	}
}

func (n *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	peer, _ := peer.FromContext(ctx)

	nc, err := makeNodeClient(version.Address)
	if err != nil {
		return nil, err
	}
	n.addPeer(nc, version)

	fmt.Printf("[Server] Received version from %s\n", peer.Addr)
	return &proto.Version{Version: n.version}, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Printf("Received tx from %s\n", peer.Addr)
	return &proto.Ack{}, nil
}

func (n *Node) addPeer(client proto.NodeClient, v *proto.Version) {
	n.plock.Lock()
	defer n.plock.Unlock()

	if (n.peers[v.Address] != nil) {
		// fmt.Printf("Skip Adding Peer %s already exists\n", v.Address)
		return
	}

	fmt.Printf("Adding peer %s version: %s\n", v.Address, v.Version)
	n.peers[v.Address] = client
}

func (n *Node) removePeer(addr string) {
	n.plock.Lock()
	defer n.plock.Unlock()

	fmt.Printf("Removing peer %s\n", addr)

	delete(n.peers, addr)
}
