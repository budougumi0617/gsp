package main

import "fmt"
import "time"

func sub1(c int) {
	fmt.Printf("share by argments", c*c)
}

func main() {
	// 引数渡し
	go sub1(10)

	// クロージャのキャプチャ渡し
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	time.Sleep(time.Second)
}
