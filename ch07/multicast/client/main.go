package main

import (
	"fmt"
	"log"
	"net"
)

// ソケットを開いて、サーバが送信するパケットを待って表示する
func main() {
	fmt.Println("Listen tick server at 224.0.0.1:9999")
	// ListenMulticastUDPを使う前に予めアドレスをパースする必要がある。
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	defer listener.Close()

	buffer := make([]byte, 1500)

	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Server   %v\n", remoteAddress)
		fmt.Printf("Now      %s\n", string(buffer[:length]))
	}
}
