package node

import (
	"fmt"
	"net"

	"github.com/hnamzian/goblock/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (n *Node) Start() error {
	opts := []grpc.ServerOption{}
	gs := grpc.NewServer(opts...)

	proto.RegisterNodeServer(gs,n)

	reflection.Register(gs)

	ln, err := net.Listen("tcp", n.addr)
	if err != nil {
		return err
	}

	fmt.Printf("Server Running on Port %s\n", n.addr)

	return gs.Serve(ln)
}
