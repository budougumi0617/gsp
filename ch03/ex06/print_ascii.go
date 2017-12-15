package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer   = strings.NewReader("COMPUTER")
	system     = strings.NewReader("SYSTEM")
	programing = strings.NewReader("PROGRAMING")
)

func main() {
	var stream io.Reader

	a := io.NewSectionReader(programing, 5, 1)
	s := io.LimitReader(system, 1)
	c := io.LimitReader(computer, 1)
	i := io.NewSectionReader(programing, 7, 1)
	i2 := io.NewSectionReader(programing, 7, 1)

	stream = io.MultiReader(a, s, c, i, i2)
	io.Copy(os.Stdout, stream)
}
