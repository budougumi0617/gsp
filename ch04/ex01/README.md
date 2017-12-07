# Q4.1 : タイマー
`time`パッケージの`time.After(duration)`により、指定した時間後に時刻データを流すチャネルが得られます。引数の`duration`には時間間隔を表す`time.Duration`型で、`10 * time.Scound`で10秒になります。  
`time.After(duration)`を使って、決まった時間を測るタイマーを作ってみましょう。

