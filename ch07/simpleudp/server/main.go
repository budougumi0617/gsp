package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server is running ad localhost:8888")
	// ListenPacketを使うと、クライアントを待たずにデータ送受信用のnet.PacketConnが即座に返される。
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// データの終了が不明なまま受信するので、バッファをつくって読み込む
	buffer := make([]byte, 1500)
	for {
		// ReadFromを使うと、接続してきた相手のアドレス情報を受け取れる
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received from %v: %v\n",
			remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"),
			remoteAddress)
		if err != nil {
			log.Fatal(err)
		}
	}
}
