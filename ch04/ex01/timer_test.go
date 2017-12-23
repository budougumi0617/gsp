package main

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s        int
		expected string
	}{
		{1, "Passed 1 seconds"},
		{0, "Passed 0 seconds"},
	}
	for _, test := range tests {
		got := Notify(test.s)
		if got != test.expected {
			t.Errorf("result = %q\nexpected = %q", got, test.expected)
		}
	}
}
