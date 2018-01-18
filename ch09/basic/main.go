package main

import "fmt"
import "io"
import "os"

// open creates new file
func open() {
	file, err := os.Create("tempfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

// read opens file
func read() {
	file, err := os.Open("tempfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func main() {
	open()
	read()
}
