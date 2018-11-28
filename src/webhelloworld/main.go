package main

import (
	"fmt"      // 文字列出力のためのpackage(formatの略)
	"net/http" // HTTPプロトコルを利用してくれるpackage(HTTPサーバーを実装するために必要な機能が提供されている)
)

// http package を使用しており,このような関数の形に書ける(handler関数)
// *http.Request型のrequestはアドレスから受け取ったパラメータを使用できる
//  L (現在のプログラムでは使用していない`/hogeとした場合にパラメータとして取得できる`) -> request.URL.Path[0:] /hogeと出力できる
func handler(writer http.ResponseWriter, request *http.Request) {
	// Helloworld をいう文字列をhttp.ResponseWriter型のwriterに入れて渡す。
	fmt.Fprintf(writer, "Hello World!")
}

func main() {
	// サーバーの中でどのようなリクエストの文字列が来たかをハンドリングする
	// つまりパスを指定してどう行った動きにするかハンドリングする
	http.HandleFunc("/", handler) // ここのhandlerはmain.go:11のhandler関数を表している。

	// サーバーを自分のPCのなかで立ち上げている
	// このおかげでgo run main.go をすると立ち上げたまま待機状態になる
	// 第一引数にport番号を渡す文字列 -> 8080 ( goでは8000や8080が多い )
	// 第二引数にnil -> DefaultServeMux が Handler として指定
	http.ListenAndServe(":8080", nil)

	// localhost(自分のPC):8080/ で表示確認できる。
	// `localhost:8080`と自分のPCの中で8080と指定した後の場所 `/`(main.go:19)を設定するとmain.go:11の関数に移動する仕組みである。
}
