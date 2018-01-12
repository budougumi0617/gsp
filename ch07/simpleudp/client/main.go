package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// MacOSだとudpではなくudp4らしい
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Sending to server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
