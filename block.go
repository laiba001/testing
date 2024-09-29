package assignment01bca

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

// Block structure
type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    Hash         string
}

// CreateHash calculates the hash of the block
func (b *Block) CreateHash() string {
    record := b.Transaction + string(b.Nonce) + b.PreviousHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

// NewBlock creates a new block
func NewBlock(transaction string, nonce int, previousHash string) *Block {
    block := &Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
    }
    block.Hash = block.CreateHash()
    return block
}

