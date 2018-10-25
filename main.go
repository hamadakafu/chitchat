// ログアウトボタンによるセッションの終了
// サインアップによる新規ユーザの登録
// テンプレートに渡す構造体の作製
// ユーザごとのtodoListの追加

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,                  // Let's Encryptの利用規約への同意
		HostPolicy: autocert.HostWhitelist("localhost"), // ドメイン名
		Cache:      autocert.DirCache("certs"),          // 証明書などを保存するフォルダ
	}
	challengeServer := &http.Server{
		Handler: certManager.HTTPHandler(nil),
		Addr:    ":8080",
	}
	go challengeServer.ListenAndServe()
	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	http.HandleFunc("/entrance", entrance)
	http.HandleFunc("/home", home)
	if err := server.ListenAndServeTLS("", ""); err != nil {
		fmt.Println(err)
	}
}

// func main() {
// 	mux := http.NewServeMux()
// 	serve := http.Server{
// 		Addr:    "127.0.0.1:8080",
// 		Handler: mux,
// 	}
// 	mux.HandleFunc("/entrance", entrance)
// 	mux.HandleFunc("/home", home)
// 	serve.ListenAndServe()
// }
