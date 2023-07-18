package types

import (
	"crypto/sha256"

	pb "github.com/golang/protobuf/proto"
	"github.com/hnamzian/goblock/internal/proto"
	"github.com/hnamzian/goblock/internal/crypto"
)

func HashBlock(block *proto.Block) ([]byte, error) {
	mb, err := pb.Marshal(block.Header)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(mb)
	return hash[:], nil
}

func SignBlock(block *proto.Block, pk crypto.PrivateKey) ([]byte, error) {
	hash, err := HashBlock(block)
	if err != nil {
		return nil, err
	}
	s := pk.Sign(hash)
	return s.Bytes(), nil
}