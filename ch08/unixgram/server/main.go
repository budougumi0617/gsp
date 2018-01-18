package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	// エラーチェックは削除（存在しなかったら存在しなかったで問題ないので不要）
	os.Remove(path)
	fmt.Println("Server is runnning at " + path)
	conn, err := net.ListenPacket("unixgram", path)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		// 返信用のクライアントのリモートアドレスが取得できる。
		// クライアント側がnet.Dialでコネクションを開いていると取得できない
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received from %v: %v\n",
			remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from sever"),
			remoteAddress)
		if err != nil {
			log.Fatal(err)
		}
	}
}
