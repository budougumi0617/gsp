package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

// UnixDomainSocketStreamServer for benchmark
func UnixDomainSocketStreamServer() {
	path := filepath.Join(os.TempDir(), "bench-unixdomainsocket-stream")
	os.Remove(path)
	listener, err := net.Listen("unix", path)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				log.Fatal(err)
			}
			_, err = httputil.DumpRequest(request, true)
			if err != nil {
				log.Fatal(err)
			}
			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 1,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
