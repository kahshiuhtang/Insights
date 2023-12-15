package network

import (
	"fmt"
	"net"
	"sync"
)

type Server struct{

}

func CreateServer() Server{
	return Server{};
}
func (s Server) Start(wg *sync.WaitGroup){
	ln, err := net.Listen("tcp", ":8081");
	if err != nil{
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}

		go handleConnection(conn, wg);
	}

}
func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
    // Close the connection when we're done
    defer conn.Close()

    // Read incoming data
    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the incoming data
    fmt.Printf("Received: %s", buf)
	wg.Done();
}