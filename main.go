package main

import (
	"fmt"
	"net"

	"github.com/hnamzian/goblock/internal/node"
	"github.com/hnamzian/goblock/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	proto.RegisterNodeServer(grpcServer, node.New())
	fmt.Printf("Server started at %s\n", ln.Addr())
	// enable reflection
	reflection.Register(grpcServer)
	grpcServer.Serve(ln)
}
