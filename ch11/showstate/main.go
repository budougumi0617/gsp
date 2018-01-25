package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // 3秒後にキャンセル
	defer cancel()
	// 第一引数のコマンドに第二引数以降のオプションを渡して実行する
	cmd := exec.CommandContext(ctx, os.Args[1], os.Args[2:]...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	select {
	// cmdが終了するのを待機する。
	case <-ctx.Done():
		fmt.Println("done:", ctx.Err())
		// default: を書くと待機せずに進む
	}
	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf("    Pid: %d\n", state.Pid())
	fmt.Printf("    System: %v\n", state.SystemTime())
	fmt.Printf("    User: %v\n", state.UserTime())
}
