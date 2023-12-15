package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/kahshiuhtang/Insights/internal/network"
)
func main() {
    fmt.Println("Hello, Go!")
	var action string;
	flag.StringVar(&action, "action", "create", "Specify an Action");
	flag.Parse();
	fmt.Println("Action", action);
	var wg sync.WaitGroup;
	serv := network.CreateServer();
	go serv.Start(&wg);
	wg.Add(1);
	wg.Add(1)
	conn, err := net.Dial("tcp", "localhost:8081");
	if err != nil{
		return;
	}
	_, err = conn.Write([]byte("Hello, server!"));
	if err != nil{
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	wg.Wait();
}