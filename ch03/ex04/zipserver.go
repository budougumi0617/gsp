package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment;filename=ascii_sample.zip")

	zw := zip.NewWriter(w)
	defer zw.Close()

	f, err := zw.Create("test")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	reader := strings.NewReader("content of file in zip")
	io.Copy(f, reader)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
