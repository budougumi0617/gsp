# 6.9 速度改善(4): パイプライニング
送受信を非同期化することでトータルの通信にかかる時間を大幅に減らす。
具体的には、レスポンスがくるまえに次から次にリクエストを多重で飛ばすことで、最終的に通信が完了するまでの時間を短くする。（ただし後方互換性がなかったり、ブラウザ・サーバ双方の実装や対応が不十分であまり使われなかった）


```bash
$ gor client/main.go
Access: 0
send:  ASCII
send:  PROGRAMING
send:  PLUS
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

```bash
$ gor server/main.go
Server is running at localhost:8888
Accept 127.0.0.1:51129
GET /?message=PROGRAMING HTTP/1.1
Host: localhost:8888
Connetion: keep-alive
User-Agent: Go-http-client/1.1


GET /?message=PLUS HTTP/1.1
Host: localhost:8888
Connetion: close
User-Agent: Go-http-client/1.1


GET /?message=ASCII HTTP/1.1
Host: localhost:8888
Connetion: keep-alive
User-Agent: Go-http-client/1.1
```
