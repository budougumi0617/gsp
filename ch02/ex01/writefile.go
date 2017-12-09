package main

import (
	"fmt"
	"log"
	"os"
)

var filename = "tmpfile.txt"

func main() {
	WriteFile(filename)
}

// WriteFile writes string in file by use of format.
func WriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "%04d, %s, %2.2f\n", 10, "test", 0.2)
}
