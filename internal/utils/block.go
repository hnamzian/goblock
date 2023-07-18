package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
	"time"

	"github.com/hnamzian/goblock/internal/proto"
)

func RandomHash() []byte {
	buf := make([]byte, 32)
	io.ReadFull(rand.Reader, buf)
	hash := sha256.Sum256(buf)
	return hash[:]
}

func GenerateRandomBlock() *proto.Block {
	return &proto.Block{
		Header: &proto.Header{
			Version:   1,
			Height:    1,
			PrevHash:  RandomHash(),
			RootHash:  RandomHash(),
			Timestamp: time.Now().UnixNano(),
		},
	}
}
