# 8.2 Unixドメインソケット版のHTTPサーバ/クライアントを実装する
Unixドメインソケットを利用して、カーネル内部で完結する高速なネットワークインタフェースを利用したHTTPサーバ/クライアントを実装する。
ソケットファイル経由で通信を行う。


# Result

```bash
$ gor client/main.go
HTTP/1.0 200 OK
Connection: close

Hello World
```

```bash
$ gor server/main.go
Server is running at /var/folders/8f/9ychhbns65s45gg8ng567nq51s7sk2/T/unixdomainsocket-sample
Accept
get / HTTP/1.1
Host: localhost:8888
User-Agent: Go-http-client/1.1

```
