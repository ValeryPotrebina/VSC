package main

import (
	"blockchain/blockchain"
	"fmt"
	"strconv"
)



func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("first")
	chain.AddBlock("second")
	chain.AddBlock("third")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		blockchain.NewProof(block)
		fmt.Println("-------------------------")

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	
	}

}
