// ログアウトボタンによるセッションの終了
// サインアップによる新規ユーザの登録
// テンプレートに渡す構造体の作製
// ユーザごとのtodoListの追加
// すべてのページでセッションの確認

package main

import (
    "crypto/tls"
    "log"
    "net/http"

    _ "github.com/lib/pq"
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS, // Let's Encryptの利用規約への同意
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
    http.HandleFunc("/entrance", entrance)
    http.HandleFunc("/home", home)
    http.HandleFunc("/createChat", createChat)
    http.HandleFunc("/successOfCreateChat", successOfCreateChat)

    if err := server.ListenAndServeTLS("", ""); err != nil {
        log.Println(err)
    }
}

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/entrance", entrance)
// 	mux.HandleFunc("/home", home)
// 	mux.HandleFunc("/createChat", createChat)
// 	mux.HandleFunc("/successOfCreateChat", successOfCreateChat)
// 	log.Fatal(http.Serve(autocert.NewListener("www.kafu-tech.xyz"), mux))
// }
// 
// func main() {
// 	mux := http.NewServeMux()
// 	serve := http.Server{
// 		Addr:    "160.16.134.28:80",
// 		Handler: mux,
// 	}
// 	mux.HandleFunc("/entrance", entrance)
// 	mux.HandleFunc("/home", home)
// 	mux.HandleFunc("/createChat", createChat)
// 	mux.HandleFunc("/successOfCreateChat", successOfCreateChat)
// 	serve.ListenAndServe()
// }
