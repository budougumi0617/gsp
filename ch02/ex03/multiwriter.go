package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func encodeJSON(w io.Writer, source map[string]string) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	return encoder.Encode(source)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"example": "encoding",
		"Hello":   "World",
	}
	// ここにコードを書く
	zipWriter := gzip.NewWriter(w)
	writers := io.MultiWriter(os.Stdout, zipWriter)
	if err := encodeJSON(writers, source); err != nil {
		log.Fatalf("%v\n", err)
	}
	zipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
