package assignment01bca

import (
    "fmt"
)

// Blockchain structure
type Blockchain struct {
    Blocks []*Block
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
    var previousHash string
    if len(bc.Blocks) != 0 {
        previousHash = bc.Blocks[len(bc.Blocks)-1].Hash
    }
    newBlock := NewBlock(transaction, nonce, previousHash)
    bc.Blocks = append(bc.Blocks, newBlock)
}

// ListBlocks prints all the blocks in the blockchain
func (bc *Blockchain) ListBlocks() {
    for _, block := range bc.Blocks {
        fmt.Printf("Transaction: %s, Nonce: %d, Previous Hash: %s, Hash: %s\n", block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
    }
}

// ChangeBlock changes the transaction of a given block
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
    if index < len(bc.Blocks) {
        bc.Blocks[index].Transaction = newTransaction
        bc.Blocks[index].Hash = bc.Blocks[index].CreateHash()
    }
}

// VerifyChain verifies the integrity of the blockchain
func (bc *Blockchain) VerifyChain() bool {
    for i := 1; i < len(bc.Blocks); i++ {
        if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].Hash {
            return false
        }
        if bc.Blocks[i].Hash != bc.Blocks[i].CreateHash() {
            return false
        }
    }
    return true
}

