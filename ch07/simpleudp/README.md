# 7.2 シンプルなUDPの実装
単純なUDPのサーバー・クライアントを実装する。


# Result

```bash
$ gor client/main.go
Sending to server
Receiving from server
Received: Hello from Server
```

```bash
$ gor server/main.go
Server is running ad localhost:8888
Received from 127.0.0.1:63587: Hello from Client
```
