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
	peers map[string]Peer

	proto.UnimplementedNodeServer
}

type Peer struct {
	Version *proto.Version
	Client proto.NodeClient
}

type Config struct {
	Addr string
}

func New(cfg *Config) *Node {
	return &Node{
		addr:    cfg.Addr,
		version: "0.0.1",
		peers:   make(map[string]Peer),
	}
}

func (n *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	peer, _ := peer.FromContext(ctx)

	nc, err := makeNodeClient(version.Address)
	if err != nil {
		return nil, err
	}
	n.addPeer(nc, version)

	// bootstrap with the peer list of the client
	n.bootstrap(version.Peers)

	fmt.Printf("[Server] Received version from %s\n", peer.Addr)
	return n.getMyVersion(), nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Printf("Received tx from %s\n", peer.Addr)
	return &proto.Ack{}, nil
}

func (n *Node) addPeer(client proto.NodeClient, v *proto.Version) {
	n.plock.Lock()
	defer n.plock.Unlock()

	if (n.peers[v.Address] != Peer{}) {
		fmt.Printf("Skip Adding Peer %s already exists\n", v.Address)
		return
	}

	fmt.Printf("My Address: %s, PeerList: %v\n", n.addr, v.Peers)

	fmt.Printf("Adding peer %s version: %s\n", v.Address, v.Version)
	n.peers[v.Address] = Peer{
		Version: v,
		Client: client,
	}
}

func (n *Node) removePeer(addr string) {
	n.plock.Lock()
	defer n.plock.Unlock()

	fmt.Printf("Removing peer %s\n", addr)

	delete(n.peers, addr)
}

func (n *Node) getMyVersion() *proto.Version {
	return &proto.Version{
		Version: n.version,
		Address: n.addr,
		Peers: n.getPeerList(),
	}
}

func (n *Node) getPeerList() []string {
	n.plock.RLock()
	defer n.plock.RUnlock()

	peers := []string{}
	for _, peer := range n.peers {
		peers = append(peers, peer.Version.Address)
	}

	return peers
}