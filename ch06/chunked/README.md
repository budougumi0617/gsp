# 6.8 速度改善(3) チャンク形式のボディー送信
チャンク形式のサーバークライアントを実装する

# Result

```bash
$ gor server/main.go
Server is running at localhost:8888
Accept 127.0.0.1:53345
GET /localhost:8888 HTTP/1.1
User-Agent: Go-http-client/1.1


```


```bash
$ gor client/main.go
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Content-Type: type/plain


 123 bytes: これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。
 117 bytes: むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、
 69 bytes: 中山さまというおとのさまが、おられたそうです。
 108 bytes: その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。
 132 bytes: ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。
 102 bytes: そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。
 ```
