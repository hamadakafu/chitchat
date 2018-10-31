// ユーザごとのtodoListの追加
// すべてのページでセッションの確認
// データベースにアクセスするための便利な構造体と関数
// success of create chatにホームに戻る方法の追加
// form の文字数制限の追加
// makeDataのなかでUserInfoを撮ってくる場合，session_idから撮ってくるようにする
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
	env := dev

	http.HandleFunc("/entrance", myhandler.Entrance)
	http.HandleFunc("/", myhandler.Login)
	http.HandleFunc("/home", myhandler.Home)
	http.HandleFunc("/signup", myhandler.SignUp)
	http.HandleFunc("/logout", myhandler.Logout)
	http.HandleFunc("/createChat", myhandler.CreateChat)
	http.HandleFunc("/successOfCreateChat", myhandler.SuccessOfCreateChat)
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

		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.Println(err)
		}
	} else if env == dev {
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Printf("What are you doing?\n")
		fmt.Printf("Please put this project in dev or pro")
	}
}
