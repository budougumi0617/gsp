package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// connで接続したサーバへ送信するリクエスト
	request, err := http.NewRequest(
		"GET", "http://localhost:8888", nil)
	if err != nil {
		log.Fatal(err)
	}

	// HttpClientに渡すときは
	request.Write(conn)
	response, err := http.ReadResponse(
		bufio.NewReader(conn), request)
	if err != nil {
		log.Fatal(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dump))
}
