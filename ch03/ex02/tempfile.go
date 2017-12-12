package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	name := "tempfile.txt"
	err := createTempFile(name)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File content size %v\n", len(data))
}

func createTempFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)
	return nil
}
