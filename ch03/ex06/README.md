# Q3.6 : ストリーム総集編
これまで紹介してきた構造体や関数を組み合わせて、ちょっとしたパズルを組み立ててみましょう。  
次の3つの文字列を3つの入力ストリーム(`io.Reader`)とし、下記に示す`main()`関数のコメント部にコードを追加して、最後の`io.Copy`で「ASCII」の文字列が出力されるようにしてみてください。

 - COMPUTER
 - SYSTEM
 - PROGRAMNG

```go
package main

import (
    "strings"
    "io"
    "os"
)

var (
    computer   = strings.NewReader("COMPUTER")
    system     = strings.NewReader("SYSTEM")
    programing = strings.NewReader("PROGRAMING")
)

func main() {
    var stream io.Reader

    // ここにioパッケージを使ったコードを書く
    io.Copy(os.Stdout, stream)
}
```

ただし次の制約を守ってください。

 - 使っていいのは`io`パッケージの内容+基本文法のみです。`io.Pipe()`を使う場合は、ブロッキングを防ぐために、次章で説明する`goroutine`を使う必要があります。
 - 文字列リテラルを使用してはいけません。
 - コメント部以外を変更してはいけません。当然、`import`するパッケージを増やしてはいけません。

ヒントとして、図3.7のモデル図を使って、ストリームの組み合わせを考えてみるといいでしょう。

# Result

```bash
$ gor print_ascii.go
ASCII%
```
