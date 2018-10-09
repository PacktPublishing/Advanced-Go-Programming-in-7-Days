package blockchain

import (
	"container/list"
	"time"
)

// Type Blockchain represents a Blockchain
type Blockchain struct {
	chain *list.List
}

func NewBlockchain(initialBlock *Block) *Blockchain  {
	chain := list.New()
	chain.PushBack(initialBlock)

	return &Blockchain{chain: chain}
}

// AddBlock adds a new Block to the back of the Blockchain with a specific transaction
func (bc *Blockchain)AddBlock(t Transaction) error {
	newBlock := &Block{}
	currentBlock := bc.GetCurrentBlock()

	newBlock.Index = currentBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = &t
	newBlock.PreviousHash = currentBlock.Hash
	bHash, err := GetBlockHash(*newBlock)
	if err != nil {
		return err
	}
	newBlock.Hash = bHash
	bc.chain.PushBack(newBlock)
	return nil
}

// GetCurrentBlock returns the tip Block
func (bc *Blockchain)GetCurrentBlock() *Block  {
	b := bc.chain.Back().Value.(*Block)
	return b
}

