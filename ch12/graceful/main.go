package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lestrrat/go-server-starter/listener"
)

func main() {
	// シグナル初期化
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// SErver::Starterからもらったソケットを確認
	listeners, err := listener.ListenAll()
	if err != nil {
		log.Fatal(err)
	}
	// Webサーバーを goroutineで起動
	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid: %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	// SIGTERMを受け取ったら終了させる
	<-signals
	server.Shutdown(context.Background())
}
