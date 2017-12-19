package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(Notify(5))
}

// Notify reuturns string after second.
func Notify(second int) string {
	var s string
	select {
	case <-time.After(time.Duration(second) * time.Second):
		s = "Passed " + strconv.Itoa(second) + " secounds"
	}
	return s
}
