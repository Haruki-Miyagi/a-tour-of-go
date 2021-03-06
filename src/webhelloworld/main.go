package main

import (
	"fmt"      // 文字列出力のためのpackage(formatの略)
	"net/http" // HTTPプロトコルを利用してくれるpackage(HTTPサーバーを実装するために必要な機能が提供されている)
)

// http package を使用しており,このような関数(handler関数)の形に書ける
// 第一引数のhttp.ResponseWriter型writerはレスポンスの書き込む
// 第二引数の*http.Request型のrequestはアドレスから受け取ったパラメータを使用できる
//  L 現在のプログラムでは使用していない(`/hoge` (HandleFunc関数の第一引数を)とした場合にパラメータとして取得できる) -> request.URL.Path[0:] で取得できる。
func handler(writer http.ResponseWriter, request *http.Request) {
	// Helloworld をいう文字列をhttp.ResponseWriter型のwriterに入れて渡す。
	fmt.Fprintln(writer, "Hello World!")
	fmt.Fprintf(writer, "Hello World!, %s", request.URL.Path[1:])
}

func main() {
	// サーバーの中でどのようなリクエストの文字列が来たかをハンドリングする
	// つまりパスを指定してどう行った動きにするかハンドリングする
	http.HandleFunc("/", handler) // ここのhandlerはmain.go:11のhandler関数を表している。

	// サーバーを自分のPCのなか(localhost)で立ち上げており、go run main.go を叩くと立ち上げたまま待機状態になる
	// 第一引数にport番号を渡す文字列 -> 8080 ( goでは8000や8080が多い )
	// 第二引数にnil -> DefaultServeMux が Handler として指定
	http.ListenAndServe(":8080", nil)

	// localhost(自分のPC):8080/ で表示確認できる。
	// `localhost:8080`と自分のPCの中で8080と指定した後の場所 `/`(main.go:19)を設定するとmain.go:11の関数に移動する仕組みである。
}
