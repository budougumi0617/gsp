package main

import "fmt"
import "sync"
import "time"

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			// ロックしてからWaitメソッドを呼ぶ
			mutex.Lock()
			defer mutex.Unlock()
			// Broadcat()が呼ばれるまで待つ
			cond.Wait()
			// 呼ばれた
			fmt.Println(name)
		}(name)
	}
	fmt.Println("よーい")
	time.Sleep(time.Second)
	fmt.Println("どん！")
	// 待っているgoroutineを一斉に起こす
	cond.Broadcast()
	time.Sleep(time.Second)
}
