package key

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	k := GenerateKey()
	assert.Equal(t, ed25519.PrivateKeySize, len(k.Bytes()), "invlaid private key size")
	assert.Equal(t, ed25519.PublicKeySize, len(k.PublicKey().Bytes()), "invalid public key size")
}

func TestNewKeyFromSeed(t *testing.T) {
	seed := "757501d7b03330d00a09d9f387c8e14d2b6f8a93c2bb5ffca5cfe87652213c20"
	k := NewKeyFromSeedString(seed)
	assert.Equal(t, ed25519.PrivateKeySize, len(k.Bytes()), "invlaid private key size")
	assert.Equal(t, ed25519.PublicKeySize, len(k.PublicKey().Bytes()), "invalid public key size")
}

func TestSignVerify(t *testing.T) {
	k := GenerateKey()
	data := []byte("hello world")
	sig := k.Sign(data)
	assert.True(t, k.PublicKey().Verify(data, sig.Bytes()), "signature verification failed")

	fk := GenerateKey()
	assert.False(t, fk.PublicKey().Verify(data, sig.Bytes()), "signature verification succeeded with wrong key")

	fdata := []byte("hello world!")
	assert.False(t, k.PublicKey().Verify(fdata, sig.Bytes()), "signature verification succeeded with wrong data")
}

func TestAddress(t *testing.T) {
	k := GenerateKey()
	assert.Equal(t, 20, len(k.PublicKey().Address().Bytes()), "invalid address size")
}
