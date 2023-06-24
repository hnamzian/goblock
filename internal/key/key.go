package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	seedLen = 32
	addressLen = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewKeyFromSeed(seed []byte) PrivateKey {
	if len(seed) != seedLen {
		panic("invalid seed length")
	}
	return PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

func NewKeyFromSeedString(seed string) PrivateKey {
	s, err := hex.DecodeString(seed)
	if err != nil {
		panic(err)
	}
	return NewKeyFromSeed(s)
}

func GenerateKey() PrivateKey {
	seed := make([]byte, seedLen)

	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	
	return NewKeyFromSeed(seed)
}

func (k PrivateKey) Bytes() []byte {
	return k.key
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: k.key.Public().(ed25519.PublicKey),
	}
}

func (k PrivateKey) Sign(data []byte) Signature {
	return Signature{
		value: ed25519.Sign(k.key, data),
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (k PublicKey) Bytes() []byte {
	return k.key
}

func (k PublicKey) Address() Address {
	return Address{
		value: k.key[12:],
	}
}

func (k PublicKey) Verify(data []byte, sig []byte) bool {
	return ed25519.Verify(k.key, data, sig)
}

type Signature struct {
	value []byte
}

func (s Signature) Bytes() []byte {
	return s.value
}

type Address struct {
	value []byte
}

func (a Address) Bytes() []byte {
	return a.value
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}