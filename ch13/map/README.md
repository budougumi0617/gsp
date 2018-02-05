# 13.7.6 sync.Map
ロックを内包し、複数の`goroutine`からアクセスされても壊れないことを保証している`map`

# Result
```bash
$ go run main.go
1: 2
hello: world
key=hello, value=world exists?=true
```
