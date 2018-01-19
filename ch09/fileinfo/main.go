package main

import "fmt"
import "os"

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(0)
	}
	// os.LStat()もある。LStatの場合、シンボリックリンクのファイルの場合、シンボリックリンクファイルの情報を得る
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf("  ファイル名： %+v\n", info.Name())
	fmt.Printf("  サイズ： %+v\n", info.Size())
	fmt.Printf("変更日時： %+v\n", info.ModTime())
	mode := info.Mode()
	fmt.Println("Mode()")
	fmt.Printf("  ディレクトリ？ %+v\n", mode.IsDir())
	fmt.Printf("  読み書き可能な通常ファイル？ = %+v\n", mode.IsRegular())
	fmt.Printf("  Unixのファイルアクセス権限ビット  %o\n", mode.Perm())
	fmt.Printf("  モードのテキスト表現  %v\n", mode.String())
}
