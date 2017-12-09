package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	source := [][]string{
		{"hello", "world"},
		{"hello", "golang"},
	}
	WriteCsv(os.Stdout, source)
}

// WriteCsv wirtes strings by CSV format.
func WriteCsv(output io.Writer, source [][]string) {
	w := csv.NewWriter(output)
	for _, line := range source {
		if err := w.Write(line); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
