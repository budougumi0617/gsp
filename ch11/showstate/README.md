# 11.5 exec.Cmdによるプロセスの起動
引数として指定された外部プログラムを実行し、それにかかった時間を表示する。
外部プログラムが3秒以内に終了しない場合は`context`によって強制終了する。

# Result
```bash
$ go run main.go sleep 2
done: context deadline exceeded
exit status 0
    Pid: 49177
    System: 501µs
    User: 583µs

$ go run main.go sleep 5
2018/01/24 09:07:41 signal: killed
exit status 1
```
