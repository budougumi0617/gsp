package main

import "log"
import "net"
import "os"
import "path/filepath"

func main() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	// エラーチェックは不要なので削除（存在しなかったらしなかったで問題ない）
	os.Remove(clientPath)
	// 受信用に自分用のソケットファイルを開いておく。
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		log.Fatal(err)
	}
	// 送信先のアドレス
	unixServerAddr, err := net.ResolveUnixAddr(
		"unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	var serverAddr net.Addr = unixServerAddr
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Sending to server")
	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Reaceived: %s\n", string(buffer[:length]))
}
