package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Start blockchain")
	bc := NewBlockchain()

	// bc.AddBlock("Send 1 BTC to MUMU")
	// bc.AddBlock("Sned 2 BTC to MOMO")

	for _, block := range bc.blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		// Validate
		pow := NewProofOfWork(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println("-----------------------------------")
	}

}
