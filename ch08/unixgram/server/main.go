package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

// ch08/domainsocket/server/main.goをデータグラム型のUNIXドメインソケット版に改造したもの。
func main() {
	// UNIX度名ソケット用のソケットファイルを準備する。
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	// 存在しなかったらしなかったで問題ないのでエラーチェックは削除
	os.Remove(path)
	listener, err := net.Listen("unixgram", path)
	if err != nil {
		log.Fatal(err)
	}
	// Ctrl + Cなどで止めてしまった場合など、サーバー側でnet.Listener.Close()を呼ばないと、ソケットファイルが残り続ける。
	defer listener.Close()
	fmt.Println("Server is running at " + path)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			// リクエストを読み込む
			request, err := http.ReadRequest(
				bufio.NewReader(conn))
			if err != nil {
				log.Fatal(err)
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(dump))
			// レスポンスを書き込む
			resp := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				// P39に説明がある。ioutil.NopCloserはダミーのClose()を持ち、io.ReadCloserのふりをする。
				Body: ioutil.NopCloser(
					strings.NewReader("Hello World\n")),
			}
			resp.Write(conn)
			// defferはつけない。defferはスコープを外れないと動かないので、
			// forループの中でdefferしたいなら、別の入れ子スコープを作らないと動かなくなる
			// https://mattn.kaoriya.net/software/lang/go/20151212021608.htm
			conn.Close()
		}()
	}
}
