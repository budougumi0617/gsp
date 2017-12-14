# Q3.4 : zipファイルをウェブサーバからダウンロード
zipファイルの出力先は単なる`io.Writer`です。そのため、2.4.4「インターネットアクセスの送信」で紹介したウェブサーバーで、zipファイルを作成して
そのままダウンロードさせるといったことも可能です。ウェブサーバーにブラウザでアクセスしたらファイルがzipダウンロードされるようにしてみましょう。  
この場合は、次のように`Content-Type`ヘッダーを使ってファイルの種類がzipファイルであることをブラウザに教えてあげる必要があります。必須ではありませんが、
ファイル名も指定できます。

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/zip")
    w.Header().Set("Content-Disposition", "attachment;filename=ascii_sample.zip")
}
```


# Result

```bash
$ unzip ~/Desktop/ascii_sample.zip
Archive:  /Users/shimizu-yoichiro/Desktop/ascii_sample.zip
  inflating: test
$ ls
README.md    test         zipserver.go
$ cat test
content of file in zip%
```
