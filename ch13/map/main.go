package main

import "fmt"
import "sync"

func main() {
	// 初期化
	smap := &sync.Map{}

	// なんでも入れられる
	smap.Store("hello", "world")
	smap.Store(1, 2)

	// 削除
	smap.Delete("test")

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
	value, ok := smap.Load("hello")
	fmt.Printf("key=%v, value=%v exists?=%v\n", "hello", value, ok)
}
