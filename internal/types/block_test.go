package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hnamzian/goblock/internal/crypto"
	"github.com/hnamzian/goblock/internal/utils"
)

func TestHashBlock(t *testing.T) {
	b := utils.GenerateRandomBlock()
	hash, err := HashBlock(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 32, len(hash), "invalid hash size")
}

func TestSignBlock(t *testing.T) {
	b := utils.GenerateRandomBlock()
	hash, err := HashBlock(b)
	if err != nil {
		t.Fatal(err)
	}
	pk := crypto.GenerateKey()
	sig, err := SignBlock(b, pk)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 64, len(sig), "invalid signature size")
	assert.True(t, pk.PublicKey().Verify(hash, sig))
}