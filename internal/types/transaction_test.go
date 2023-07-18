package types

import (
	"testing"

	"github.com/hnamzian/goblock/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHash(t *testing.T) {
	tx := utils.GenerateRandomTransaction()
	
	hash, err := HashTransaction(tx)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 32, len(hash))
}

func TestVerifyTransction(t *testing.T) {
	tx := utils.GenerateRandomTransaction()

	v, err := VerifyTxSignature(tx)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, v)
}