package main

import (
    "fmt"
    "github.com/YourGithubUsername/assignment01bca"
)

func main() {
    bc := &assignment01bca.Blockchain{}
    
    bc.AddBlock("alice to bob", 1)
    bc.AddBlock("bob to charlie", 2)
    bc.AddBlock("charlie to dave", 3)

    fmt.Println("Initial Blockchain:")
    bc.ListBlocks()

    bc.ChangeBlock(1, "bob to eve")
    fmt.Println("\nBlockchain after changing transaction in block 2:")
    bc.ListBlocks()

    fmt.Printf("\nBlockchain verification: %v\n", bc.VerifyChain())
}

