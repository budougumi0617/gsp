# 6.6.1 Keep-Alive対応のHTTPサーバー
`tcp01`のHTTPサーバーにKeep-Aliveを実装する。HTTP/1.1で採用されたKeep-Alive規格はしばらくの間TCP接続のコネクションを維持して使いまわす。


# Result

```bash
$ gor server/main.go
Server is running at localhost:8888
Accept 127.0.0.1:57410
PORT / HTTP/1.1
Host: localhost:8888
Content-Length: 5
User-Agent: Go-http-client/1.1

ASCII
PORT / HTTP/1.1
Host: localhost:8888
Content-Length: 10
User-Agent: Go-http-client/1.1

PROGRAMING
PORT / HTTP/1.1
Host: localhost:8888
Content-Length: 4
User-Agent: Go-http-client/1.1

PLUS
```

```bash
$ gor client/main.go
Access: 0
HTTP/1.1 200 OK
Content-Length: 12

Hello World

HTTP/1.1 200 OK
Content-Length: 12

Hello World

HTTP/1.1 200 OK
Content-Length: 12

Hello World
```