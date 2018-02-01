# 13.7.2 sync.WaitGroup
`Wait()`したとき、`Add()`された数だけ`Done()`を待機する。
`sync.WaitGroup`構造体は値コピーすると正しく動作しない。変数宣言だけで利用できる
