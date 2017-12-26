package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			for {
				// タイムアウトを設定
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				// リクエストを埋め込む
				request, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					// タイムアウトもしくはソケットクローズ時は終了
					neterr, ok := err.(net.Error) // ダウンキャスト
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				// リクエストを表示
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"
				// レスポンスを書き込む
				// HTTP/1.1かつ、ContentLengthの設定が必要

				resp := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					// P39に説明がある。ioutil.NopCloserはダミーのClose()を持ち、io.ReadCloserのふりをする。
					Body: ioutil.NopCloser(
						strings.NewReader(content)),
				}
				resp.Write(conn)
			}
		}()
	}
}
