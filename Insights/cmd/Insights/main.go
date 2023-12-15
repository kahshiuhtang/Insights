package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kahshiuhtang/Insights/internal/network"
)
func main() {
	var serverName string;
	var serverAddr string;
	flag.StringVar(&serverName, "name", "None", "Specify a server name");
	flag.StringVar(&serverAddr, "addr", "None", "Specify a server address")
	flag.Parse();

	if serverName == "None"{
		fmt.Println("[Main]: Server must have a name.")
		return;
	}
	if serverAddr == "None"{
		fmt.Println("[Main]: Server must have a address.")
		return;
	}
	msgChan := make(chan string);
	serv := network.CreateServer(serverName, serverAddr, msgChan);
	go serv.Start();

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		lowercaseInput :=  strings.ToLower(input);
		if lowercaseInput == "exit"{
			fmt.Println("Exiting...");
			break;
		}else if lowercaseInput == "send"{
			
			fmt.Println("");
		}
	}
}