package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	conn, err := net.Dial("unix",
		filepath.Join(os.TempDir(), "unixdomainsocket-sample"))
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest(
		"get", "http://localhost:8888", nil)
	if err != nil {
		log.Fatal(err)
	}
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
	fmt.Printf(string(dump))
}
