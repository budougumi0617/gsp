package main

import "archive/zip"
import "io"

import "log"
import "os"
import "strings"

func main() {
	name := "temp.zip"
	zipfile, err := os.Create(name)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer zipfile.Close()

	zw := zip.NewWriter(zipfile)
	zw.Close()

	f, err := zw.Create("test")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	reader := strings.NewReader("content of file in zip")
	io.Copy(f, reader)
}
