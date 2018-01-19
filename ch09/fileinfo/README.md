# 9.2.4 ファイルの属性の取得
`os.Stat()`を利用してファイルの属性を取得する


```bash
gor main.go ../../README.md
FileInfo
  ファイル名： README.md
  サイズ： 575
変更日時： 2017-12-10 08:21:21 +0900 JST
Mode()
  ディレクトリ？ false
  読み書き可能な通常ファイル？ = true
  Unixのファイルアクセス権限ビット  644
  モードのテキスト表現  -rw-r--r--
```
