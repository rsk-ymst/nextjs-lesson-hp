package main

import (
    "fmt"
    "net/http"
)

// ChitChatのデータモデル
// ユーザー：フォーラムのユーザ情報
// セッション：ユーザーが現在ログインしているセッション
// スレッド：フォーラムのスレッド
// ポスト：スレッド内の投稿

func handler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:]) // ブラウザに表示
}

func main() {
    mux := http.NewServeMux() // マルチプレクサを生成
    files := http.FileServer(http.Dir("/public")) // リクエスト
    mux.Handle("/static/", http.StripPrefix("/static/", files))

    mux.HandleFunc("/", index)
    mux.HandleFunc("/err", err)
    mux.HandleFunc("/logout", logout)
    mux.HandleFunc("/signup", signup)

    server := &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: mux,
    }

    server.ListenAndServe()
}
