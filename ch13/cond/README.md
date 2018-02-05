# 13.7.4 sync.Cond
`sync.Cond`を使えば`groutine`の数がゼロであっても複数であっても通知することができる。
チャネルの場合は、待っているすべての`goroutine`に通知するとしたらクローズするしかないため、一度きりの通知にしか使えない。

# Result

```bash
$ go run  main.go
よーい
どん！
C
A
B
```
