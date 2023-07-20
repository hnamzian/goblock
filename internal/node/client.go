package node

import (
	"github.com/hnamzian/goblock/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func makeNodeClient(addr string) (proto.NodeClient, error) {
	client, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	cc := proto.NewNodeClient(client)

	return cc, nil
}
