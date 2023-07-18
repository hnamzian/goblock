package types

import (
	"crypto/sha256"

	pb "github.com/golang/protobuf/proto"
	"github.com/hnamzian/goblock/internal/crypto"
	"github.com/hnamzian/goblock/internal/proto"
)

func GetUnsignedTx(tx *proto.Transaction) *proto.Transaction {
	rawTx := proto.Transaction{
		Version: tx.Version,
		Outputs: tx.Outputs,
	}
	for _, in := range tx.Inputs {
		rawTx.Inputs = append(rawTx.Inputs, &proto.TxInput{
			PrevTxHash: in.PrevTxHash,
			PrevTxIndex: in.PrevTxIndex,
			Pubkey: in.Pubkey,
		})
	}
	return &rawTx
}

func HashTransaction(tx *proto.Transaction) ([]byte, error) {
	mt, err := pb.Marshal(tx)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(mt)
	return hash[:], nil
}

func VerifyTxSignature(tx *proto.Transaction) (bool, error) {
	rawTx := GetUnsignedTx(tx)
	hash, err := HashTransaction(rawTx)
	if err != nil {
		return false, err
	}

	for _, in := range tx.Inputs {
		pbkey := crypto.PublicKeyFromBytes([]byte(in.Pubkey))
		sig := []byte(in.Signature)
		v := pbkey.Verify(hash, sig)
		if !v {
			return false, nil
		}
	}
	return true, nil
}