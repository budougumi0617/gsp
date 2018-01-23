package main

import "fmt"
import "os"
import "path/filepath"
import "strings"

var imageSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tif":  true,
	".eps":  true,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(`Find images

Usage:
    %s [path to fine]
`, os.Args[0])
		return
	}
	root := os.Args[1]

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if info.Name() == "_build" { // Sphinxの出力ディレクトリ名
					return filepath.SkipDir
				}
				return nil
			}
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if imageSuffix[ext] {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return nil
				}
				fmt.Printf("%s\n", rel)
			}
			return nil
		})
	if err != nil {
		fmt.Println(1, err) // 1はなんだろう？
	}
}
