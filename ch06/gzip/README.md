# 6.7 速度改善(2) 圧縮
HTTPサーバとクライアントに一般的なブラウザでも使われているgzip圧縮を実装する。



# Result

```bash
$ gor server/main.go
Server is running at localhost:8888
Accept 127.0.0.1:53197
PORT / HTTP/1.1
Host: localhost:8888
Accept-Encoding: gzip
Content-Length: 5
User-Agent: Go-http-client/1.1

ASCII
PORT / HTTP/1.1
Host: localhost:8888
Accept-Encoding: gzip
Content-Length: 10
User-Agent: Go-http-client/1.1

PROGRAMING
PORT / HTTP/1.1
Host: localhost:8888
Accept-Encoding: gzip
Content-Length: 4
User-Agent: Go-http-client/1.1

PLUS
```


```bash
$ gor client/main.go
Access: 0
HTTP/1.1 200 OK
Content-Length: 22


Hello World (gzipped)
HTTP/1.1 200 OK
Content-Length: 22


Hello World (gzipped)
HTTP/1.1 200 OK
Content-Length: 22


Hello World (gzipped)
```
