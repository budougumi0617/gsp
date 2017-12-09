package main

import (
	"io/ioutil"
	"testing"
)

func TestWriteFile(t *testing.T) {
	var tests = []struct {
		filename, expected string
	}{
		{"tempfile.txt", "0010, test, 0.20\n"},
	}
	for _, test := range tests {
		WriteFile(test.filename)
		got, err := ioutil.ReadFile(test.filename)
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		if string(got) != test.expected {
			t.Errorf("result = %q\n, expected = %q", got, test.expected)
		}
	}
}
