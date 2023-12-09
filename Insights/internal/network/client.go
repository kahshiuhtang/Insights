package network

import (
	"bufio"
	"context"
	"fmt"
	"os"

	lib2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	net "github.com/libp2p/go-libp2p/core/network"
)
func handleStream(stream net.Stream){
	rw:= bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go readData(rw);
	go writeData(rw);
}
func readData(rw *bufio.ReadWriter){
	for{
		str, err := rw.ReadString('\n');
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}
		if str != "\n"{
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}
	}
}
func writeData(rw *bufio.ReadWriter){
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n');

		if err != nil {
			fmt.Println("Error reading from STDIN");
			panic(err)
		}
		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData));
		if err != nil {
			fmt.Println("Error writing to buffer");
			panic(err);
		}
		err = rw.Flush()
		if err != nil{
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}
}
func CreateClient(){
	host, err := lib2p.New()
	if err != nil {
		fmt.Println("Error creating new");
		panic(err)
	}
	host.SetStreamHandler("/conn/1.1.0", handleStream);

	ctx := context.Background()
	dht, err := dht.New(ctx, host);
	if err != nil{
		panic(err)
	}
	if err = dht.Bootstrap(ctx); err != nil{
		panic(err)
	}
	}