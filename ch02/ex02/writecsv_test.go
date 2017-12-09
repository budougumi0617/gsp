package main

import (
	"bytes"
	"testing"
)

func TestWriteCsv(t *testing.T) {
	var tests = []struct {
		output   *bytes.Buffer
		source   [][]string
		expected string
	}{
		{&bytes.Buffer{},
			[][]string{{"hello", "world"}, {"good", "bye"}},
			"hello,world\ngood,bye\n"},
	}
	for _, test := range tests {
		WriteCsv(test.output, test.source)
		got := test.output.String()
		if string(got) != test.expected {
			t.Errorf("result = %q\n, expected = %q", got, test.expected)
		}
	}
}
