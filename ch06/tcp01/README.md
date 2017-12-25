# 6.5.1 TCPソケットを使ったHTTPサーバー
Go言語のソケットを使ってHTTP/1.0相当の送受信を実現してみましょう。


# Result log

```bash
Accept 127.0.0.1:64012
GET / HTTP/1.1
Host: localhost:8888
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: ja,en-US;q=0.9,en;q=0.8
Cache-Control: max-age=0
Connection: keep-alive
Cookie: _app_session=NUVYeGh...54414.1513840574.1514097775.13
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36


Accept 127.0.0.1:64014
GET /favicon.ico HTTP/1.1
Host: localhost:8888
Accept: image/webp,image/apng,image/*,*/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: ja,en-US;q=0.9,en;q=0.8
Cache-Control: no-cache
Connection: keep-alive
Cookie: _app_session=NUVYeGhCR...54414.1513840574.1514097775.13
Pragma: no-cache
Referer: http://localhost:8888/
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36
```
