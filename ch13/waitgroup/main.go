package main

import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup

	// ジョブ数をあらかじめ登録しておく
	wg.Add(2)

	go func() {
		// 非同期で仕事をする（1)
		fmt.Println("Job 1")
		// Doneで完了通知
		wg.Done()
	}()

	go func() {
		// 非同期で仕事をする（2)
		fmt.Println("Job 2")
		// Doneで完了通知
		wg.Done()
	}()

	// 全ての処理が終わるのを待機する
	wg.Wait()
	fmt.Println("終了")
}
