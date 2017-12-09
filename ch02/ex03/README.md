# Q2.3: gzipされたJSON出力をしながら、標準出力にログを出力
JSONをgzip化してクライアントに返すウェブサービスを開発しているとします。本章で説明したように`http.ResponseWriter`を指定してJSONエンコーダーで変換をかけるとJSONが生成できますが、そのままではどのようなレスポンスをクライアントに返したのかログを残すことができません。`os.Stdout`に出力するとログが残るものとして、JSONの文字列変換、gzip圧縮を行いながら圧縮前の出力を標準出力にもだすように、`io.MultiWriter`を使ってみましょう。gzip出力の最後に`Flush()`が必要な点に注意してください。

```go
package main

import (
    "compress/gzip"
    "encoding/json"
    "io"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Encoding", "gzip")
   w.Header().Set("Content-Type", "application/json")
   // json化する元のデータ
   source := map[string]string{
       "Hello": "World",
   }
   // ここにコードを書く
}

func main() {
  http.HandleFunc("/", handler)
  http.listenAndServe(":8080", nil)
}
```
