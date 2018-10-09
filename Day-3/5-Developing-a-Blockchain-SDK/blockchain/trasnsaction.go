package blockchain

import (
	"crypto/sha256"
	"bytes"
	"encoding/binary"
)

type Transaction struct {
	Hash []byte
	Payload []byte
}

func NewTransaction(payload []byte) *Transaction  {
	tx := Transaction{Payload: payload}
	tx.Hash, _ = GetTransactionHash(tx)

	return &tx
}

// Get the transaction hashcode that is calculated from the payload
func GetTransactionHash(tx Transaction) ([]byte, error)  {
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(txBytes)

	return h.Sum(nil), nil
}

func (tx *Transaction)MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, tx.Payload)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


