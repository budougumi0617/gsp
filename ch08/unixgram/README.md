# 8.2.4 データグラム型のUnixドメインソケット
UDP相当の使い方ができるデータグラム型のUnixドメインソケットの実装をする。


# Result

```bash
gor server/main.go
Server is runnning at /var/folders/8f/9ychhbns65s45gg8ng567nq51s7sk2/T/unixdomainsocket-server
Received from /var/folders/8f/9ychhbns65s45gg8ng567nq51s7sk2/T/unixdomainsocket-client: Hello from Client
```

```bash
gor client/main.go
2018/01/18 08:55:49 Sending to server
2018/01/18 08:55:49 Receiving from server
2018/01/18 08:55:49 Reaceived: Hello from sever
```
