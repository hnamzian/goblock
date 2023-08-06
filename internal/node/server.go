package node

import (
	"fmt"
	"net"
	"time"

	"github.com/hnamzian/goblock/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (n *Node) Start(staticNodes []string) error {
	opts := []grpc.ServerOption{}
	gs := grpc.NewServer(opts...)

	proto.RegisterNodeServer(gs,n)

	reflection.Register(gs)

	ln, err := net.Listen("tcp", n.addr)
	if err != nil {
		return err
	}

	fmt.Printf("Server Running on Port %s\n", n.addr)

	go func() {
		n.bootstrap(staticNodes)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		n.monitor()	
	}()
	

	return gs.Serve(ln)
}
