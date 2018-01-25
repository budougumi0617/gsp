package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	// サイズが１より大きいチャネルを作成
	siganls := make(chan os.Signal, 1)
	// 最初のチャネル以降の引数は、可変長引数で任意の数のシグナルを設定可能
	signal.Notify(siganls, syscall.SIGINT, syscall.SIGTERM)

	// シグナルがクルマで待機
	s := <-siganls

	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}
