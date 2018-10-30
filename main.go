// ログアウトボタンによるセッションの終了
// テンプレートに渡す構造体の作製
// ユーザごとのtodoListの追加
// すべてのページでセッションの確認
// postgres 開発環境と本番環境の統一　ユーザの権限
// データベースにアクセスするための便利な構造体と関数
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"./myhandler"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	dev := "dev"
	pro := "pro"
	// 環境設定
	env := pro
	if env == pro {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,                          // Let's Encryptの利用規約への同意
			HostPolicy: autocert.HostWhitelist("www.kafu-tech.xyz"), // ドメイン名
			Cache:      autocert.DirCache("/root/cache"),            // 証明書などを保存するフォルダ
		}

		// http-01 Challenge(ドメインの所有確認)、HTTPSへのリダイレクト用のサーバー
		challengeServer := &http.Server{
			Handler: certManager.HTTPHandler(nil),
			Addr:    ":80",
		}

		go challengeServer.ListenAndServe()

		server := &http.Server{
			Addr: ":443",
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
		http.HandleFunc("/entrance", myhandler.Entrance)
		http.HandleFunc("/home", myhandler.Home)
		http.HandleFunc("/signup", myhandler.SignUp)
		http.HandleFunc("/logout", myhandler.Logout)
		//http.HandleFunc("/createChat", myhandler.CreateChat)
		//http.HandleFunc("/successOfCreateChat", myhandler.successOfCreateChat)

		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.Println(err)
		}
	} else if env == dev {
		mux := http.NewServeMux()
		serve := http.Server{
			Addr:    "127.0.0.1:8080",
			Handler: mux,
		}
		mux.HandleFunc("/entrance", myhandler.Entrance)
		mux.HandleFunc("/home", myhandler.Home)
		mux.HandleFunc("/signup", myhandler.SignUp)
		mux.HandleFunc("/logout", myhandler.Logout)
		// mux.HandleFunc("/createChat", myhandler.CreateChat)
		// mux.HandleFunc("/successOfCreateChat", myhandler.SuccessOfCreateChat)
		serve.ListenAndServe()
	} else {
		fmt.Printf("What are you doing?\n")
		fmt.Printf("Please put this project in dev or pro")
	}
}
