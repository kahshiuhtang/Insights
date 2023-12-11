package network

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"

	lib2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	net "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
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
	config, err := ParseFlags();
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

	//Bootstrap DHT, spawns background thread that refreshs peer table ~5 minutes
	if err = dht.Bootstrap(ctx); err != nil{
		panic(err)
	}

	//Connect to bootstrap nodes, tell us about other nodes in network
	var wg sync.WaitGroup;
	for _, peerAddr := range config.BootstrapPeers{
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr);
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err :=  host.Connect(ctx, *peerinfo); err != nil{
				fmt.Println(err);
			}else{
				fmt.Println("Connection established with bootstrap node")
			}
		}()
	}
	wg.Wait();
	
	//Rendezvous for others
	routingDiscovery := drouting.NewRoutingDiscovery(dht)
	dutil.Advertise(ctx, routingDiscovery, config.RendezvousString)

	// Finding other hosts who have also announced
	peerChan, err := routingDiscovery.FindPeers(ctx, config.RendezvousString);

	if err != nil{
		panic(err);
	}
	for peer := range peerChan {
		if peer.ID == host.ID(){
			continue
		}
		stream, err := host.NewStream(ctx, peer.ID, protocol.ID(config.ProtocolID))

		if err != nil {
			continue;
		}else{
			rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream));
			go writeData(rw);
			go readData(rw);
		}
		
	}
}