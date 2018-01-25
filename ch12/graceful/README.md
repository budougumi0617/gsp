# 12.5 Server::Starter対応のGracefull Shutdownをするサーバーの簡易実装
`Server::Starter`提供のソケットを使って起動して`http.Server.Shutdown()`メソッドを使ってGraceful Shutdownを実行する。


```bash
$ go get github.com/lestrrat/go-server-starter/cmd/start_server
$ go build -o server main.go
$ start_server --port 8080 --pid-file app.pid -- ./server
$ start_server --port 8080 --pid-file app.pid -- ./server
starting new worker 82738
---------------------------------------Other terminal
$ cat app.pid
82737%
$ kill -HUP `cat app.pid`
---------------------------------------
received HUP (num_old_workers=TODO)
spawning a new worker (num_old_workers=TODO)
starting new worker 86037
new worker is now running, sending TERM to old workers:82738
sleep 0 secs
killing old workers
old worker 82738 died, status:0
^Creceived INT, sending TERM to all workers:86037
worker 86037 died, status:2
exiting
```
