package multichan

import (
	"net"
	"net/http"
)

// 終了した順に書き出し
// チャネルに結果が投入された順に処理される
func writeToConnSimple(responses chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for response := range responses {
		response.Write(conn)
	}
}

// 開始した順に書き出し
// チャネルにチャネルを入れた（開始した）順に処理される
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
	}
}
