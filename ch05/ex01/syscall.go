package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	t, err := os.Create("syscall.trace")
	if err != nil {
		panic(err)
	}
	defer t.Close()

	err = trace.Start(t)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	f, err := os.Create("tempfile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 100; i++ {
		fmt.Fprintf(f, "line number: %d\n", i)
	}
}
