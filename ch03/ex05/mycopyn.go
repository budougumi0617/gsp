package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// CopyN implementes io.CopyN.
func CopyN(dest io.Writer, src io.Reader, length int) (int, error) {
	length64 := int64(length)
	lr := io.LimitReader(src, length64)
	written, err := io.Copy(dest, lr)
	if written == length64 {
		return int(written), nil
	}
	if written < length64 && err == nil {
		return int(written), io.EOF
	}
	return int(written), err
}

func main() {
	s := "test test test"
	buf := strings.NewReader(s)
	fmt.Printf("origin string = %v\n", s)
	n, err := CopyN(os.Stdout, buf, 3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\ncopied size = %d\n", n)
}
