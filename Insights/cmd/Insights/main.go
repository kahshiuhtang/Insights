package main

import (
	"fmt"

	"github.com/kahshiuhtang/Insights/internal/blockchain"
	"github.com/kahshiuhtang/Insights/internal/network"
)

func main() {
    fmt.Println("Hello, Go!")
	blockchain.CreateBlockchain(5);
	network.CreateClient();
}