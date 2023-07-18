package utils

import (
	"crypto/sha256"

	pb "github.com/golang/protobuf/proto"
	"github.com/hnamzian/goblock/internal/crypto"
	"github.com/hnamzian/goblock/internal/proto"
)

func SignTx(tx *proto.Transaction, pk crypto.PrivateKey) ([]byte, error) {
	txBuf, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(txBuf)
	
	sigTx := pk.Sign(hash[:])

	return sigTx.Bytes(), nil
}

func GenerateRandomTransaction() *proto.Transaction {
	fromPK := crypto.GenerateKey()
	toPK := crypto.GenerateKey()

	txInputs := []*proto.TxInput{
		{
			PrevTxHash: RandomHash(),
			PrevTxIndex: 0,
			Pubkey: fromPK.PublicKey().Bytes(),
		},
	}
	txOutputs := []*proto.TxOutput{
		{
			Value: 100,
			Address: toPK.PublicKey().Address().Bytes(),
		},
	}
	tx := &proto.Transaction{
		Version: 1,
		Inputs: txInputs,
		Outputs: txOutputs,
	}

	sigTx, err := SignTx(tx, fromPK)
	if err != nil {
		panic(err)
	}

	for _, txInput := range txInputs {
		txInput.Signature = sigTx
	}
	
	return tx
}