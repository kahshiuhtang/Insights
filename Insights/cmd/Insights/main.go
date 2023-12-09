package main

import (
	"fmt"
	"insights/internal/blockchain"
)

func main() {
    fmt.Println("Hello, Go!")
	blockchain.CreateBlockchain(5);
}