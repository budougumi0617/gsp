package main

import (
	"io/ioutil"
	"testing"
)

func TestCreateTempFile(t *testing.T) {
	name := "tempfile.txt"
	err := createTempFile(name)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if 1024 != len(data) {
		t.Fatalf("Incorrect file size %d\n", len(data))
	}
}
