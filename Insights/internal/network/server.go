package network

import (
	"fmt"
	"net"
)

type Server struct{
	msgChan chan string 
	serverName string
	port string
	connection net.Conn
}

func CreateServer(name string, port string, msgChan chan string) Server{
	return Server{msgChan: msgChan, serverName: name, port: port};
}
func (s Server) Start(){
	ln, err := net.Listen("tcp", s.port);
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

		go handleConnection(conn);
	}

}
func handleConnection(conn net.Conn) {
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
}
func (s Server) Shutdown() bool{
	close(s.msgChan);
	return true;
}
func (s Server) SendMessage(address string, msg string) {
	// "localhost:8081"
	conn, err := net.Dial("tcp", address);
	if err != nil{
		return;
	}
	_, err = conn.Write([]byte(msg));
	if err != nil{
		return
	}
}