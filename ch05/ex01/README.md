# Q5.1 : システムコールの確認
本章の最後にシステムコールやWin32 APIの呼び出しをモニターするようなツールを紹介しました。ファイルの入出力のサンプルコードを作成し、どのようなシステムコールやAPIが呼ばれているか確認してみましょう。



# Result

```bash
$ go tool trace syscall.trace
```

Macbook Pro2015で以下のメソッドで`Syscall`が呼ばれていることを確認。

https://github.com/golang/go/blob/master/src/syscall/zsyscall_darwin_amd64.go#L1321

```
func write(fd int, p []byte) (n int, err error) {
	var _p0 unsafe.Pointer
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
```
