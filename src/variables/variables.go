package main

import (
	"fmt"
	"reflect"
)

// boolは boolean型
var c, python, java bool

func main() {
	// int数字を扱う型であり、変数名 i を宣言する
	var i int
	fmt.Println(i, c, python, java) // OK ! (初期値のチェック)
	fmt.Println("--- 変数書き換え ---")
	// 変数代入 i に 100を代入する
	i = 100
	fmt.Println(i) // OK!
	// 書き換え可能
	i = 200
	fmt.Println(i) // OK!

	fmt.Println("--- int型のチェック ---")
	var s = 10000
	fmt.Println(s)                 // OK!
	fmt.Println(reflect.TypeOf(s)) // OK! #=> int型
}
