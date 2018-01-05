package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn
	//リトライ用にループで全体を囲う
	for {
		var err error
		// まだコネクションを張ってない / エラーでリトライ
		if conn == nil {
			// Dialから行ってconnを初期化
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Access: %d\n", current)
		}
		// POSTで文字列を送るリクエストを作成
		request, err := http.NewRequest(
			"PORT",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]))
		if err != nil {
			log.Fatal(err)
		}
		// このクライアントはgzipを処理できます
		request.Header.Set("Accept-Encoding", "gzip")
		err = request.Write(conn)
		if err != nil {
			log.Fatal(err)
		}
		// サーバーから読み込む。タイムアウトはここでエラーになるのでリトライ
		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		// 結果を表示。第二引数をfalseにして、圧縮されている可能性があるBodyは無視する。
		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()

		// 圧縮されている場合は復元する
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}
		// 全部送信完了していれば終了
		current++
		if current == len(sendMessages) {
			break
		}
	}
	// 全て終わったらCloseする
	conn.Close()
}
