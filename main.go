// ログアウトボタンによるセッションの終了
// サインアップによる新規ユーザの登録
// テンプレートに渡す構造体の作製
// ユーザごとのtodoListの追加
// すべてのページでセッションの確認

package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

// func main() {
//     mux := http.NewServeMux()
// 	mux.HandleFunc("/entrance", entrance)
// 	mux.HandleFunc("/home", home)
//     log.Fatal(http.Serve(autocert.NewListener("www.kafu-tech.xyz"), mux))
// }

func main() {
	mux := http.NewServeMux()
	serve := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/entrance", entrance)
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/createChat", createChat)
	mux.HandleFunc("/successOfCreateChat", successOfCreateChat)
	serve.ListenAndServe()
}
