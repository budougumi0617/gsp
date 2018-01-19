package main

import "fmt"
import "io"
import "os"

// 新規作成
func open() {
	file, err := os.Create("tempfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

// 読み込み。ファイルオープン時のオプションが違う
func read() {
	file, err := os.Open("tempfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

// 追記
func append() {
	file, err := os.OpenFile("tempfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Apended content\n")
}

// フォルダを1個だけ作成
func mkdir() {
	os.Mkdir("setting", 0755)
}

// 深いフォルダを1回で作成
func mkdirAll() {
	os.MkdirAll("setting/myapp/networksettings", 0755)
}

// ファイルや空のディレクトリの削除
func remove() {
	os.Remove("tempfile.txt")
}

// ディレクトリを中身ごと削除
func removeAll() {
	os.RemoveAll("setting")
}

// 先頭100バイトで切る
func truncate() {
	os.Truncate("tempfile.txt", 100)
}

// Truncateメソッドを利用する場合
func mTruncate(name string) {
	file, _ := os.Open(name)
	file.Truncate(100)
}

// リネーム（移動。移動先にディレクトリは指定できない）
func rename(old, new string) {
	os.Rename(old, new)
}

// 移動先のデバイスやドライブが異なる場合にはファイルを開いてコピーする必要がある
func move(old, new string) {
	oldFile, err := os.Open(old)
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create(new)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		panic(err)
	}
	oldFile.Close()
	os.Remove(old)
}

func main() {
	open()
	read()
}
