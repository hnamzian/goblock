package node

import (
	"context"
	"fmt"
	"time"

	"github.com/hnamzian/goblock/internal/proto"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (n *Node) monitor() {
	// n.plock.Lock()
	// defer n.plock.Unlock()
	for {
		peerPool := []string{}
		for addr, peer := range n.peers {
			myVersion := n.getMyVersion()
			v, err := peer.Client.Handshake(context.Background(), myVersion)
			if err != nil {
				fmt.Printf("Error to handshake with %s: %s\n", addr, err)
				n.removePeer(addr)
			}
			// import peers of client in a pool to bootstrap nodes later
			peerPool = append(peerPool, v.Peers...)
		}
		// now bootstrap nodes with all peers in the pool
		if (len(peerPool) > 0) {
			n.bootstrap(peerPool)
		}

		time.Sleep(5 * time.Second)
	}
}

func (n *Node) bootstrap(staticNodes []string) {
	nodes := staticNodes

	// remove duplicate nodes and already connected nodes and self from the list
	nodes = n.getStrangeNodes(nodes)

	if len(nodes) == 0 {
		return
	}
	fmt.Printf("Bootstrapping with %d nodes\n", len(nodes))

	// n.plock.Lock()
	// defer n.plock.Unlock()
	for len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			addr := nodes[i]
			c, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Printf("Error to dial %s: %s\n", addr, err)
				continue
			}

			nc := proto.NewNodeClient(c)

			v, err := nc.Handshake(context.Background(), &proto.Version{Version: n.version, Address: n.addr, Height: 0})
			if err != nil {
				fmt.Printf("Error to handshake with %s: %s\n", addr, err)
				continue
			}

			n.addPeer(nc, v)
			if (i == len(nodes) - 1) {
				nodes = nodes[:i]
			} else {
				nodes = slices.Delete(nodes, i, i+1)
			}
			i--
		}

		fmt.Printf("%d nodes remaining\n", len(nodes))
		if (len(nodes) > 0) {
			fmt.Println("sleeping...")
			time.Sleep(5 * time.Second)
		}
	}

	fmt.Println("Bootstrap done")
}

func (n *Node) getStrangeNodes(staticNodes []string) []string {
	nodes := staticNodes
	strangeNods := []string{}

	// remove duplicate nodes
	nodes = removeDuplicateStr(nodes)

	// append all nodes except self and already connected nodes
	for _, node := range nodes {
		if node != n.addr && (n.peers[node] == Peer{}) {
			strangeNods = append(strangeNods, node)
		}
	}

	return strangeNods
}

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}