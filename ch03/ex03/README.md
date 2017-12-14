# Q3.3 : zipファイルの書き込み
OSのデバイスにリンクされた`io.Writer`や`io.Reader`は、1つのファイルやデバイスと1対1に対応しています。Go言語が提供するライブラリには、1つのファイルで複数の`io.Writer`や`io.Reader`に対応しているものもあります。
複数ファイルを格納するアーカイブフォーマットである`tar`や`zip`ファイルや、インターネットのマルチパート形式(ブラウザのフォームによって作られるデータやファイルを複数格納するデータ構造)をサポートする`mime/multipart`パッケージの構造体は、中に格納されるひとつひとつの要素が`io.Writer`や`io.ReadCloser`になっています。  
`archive/zip`パッケージを使ってzipファイルを作成してみましょう。出力先のファイルの`Writer`(以下のコードの`file`)をまず作って、それを`zip.NewWriter()`関数に渡すと、zipファイルの書き込み用の構造体が出来ます。最後に`Close()`を確実に呼ぶ必要がありますが、これにはGo言語の`defer`という機能を使って次のようにすればよいでしょう。

```go
zipWriter := zip.NewWriter()
defer zipWriter.Close()
```

この構造体そのものは`io.Writer`ではありませんが、`Create()`メソッドを呼ぶと、個別のファイルを書き込むための`io.Writer`が返ってきます。

```go
writer, err := zipWriter.Create("newfile.txt")
```

上記の例では、`newfile.txt`という実際のファイルが、最初に作った出力先ファイル`file`へと圧縮されます。では、実際のファイルではなく、文字列`strings.Reader`を使ってzipファイルを作成するにはどうすればいいでしょうか。考えてみてください。


# Result

```bash
$ gor zipstrings.go
$ unzip temp.zip
Archive:  temp.zip
  inflating: test
$ cat test
content of file in zip%
```
