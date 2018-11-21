package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
)

// boolは boolean型
var c, python, java bool

var (
	ToBe bool = false
	// 絶対値-> uint , int -> 整数
	// float -> 浮動小数点
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	s      string
)

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

	fmt.Println("--- typeの変換 ---")
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	fmt.Println("--- basic type の確認 ---")
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)

	fmt.Println("string型の初期値")
	fmt.Printf("%q\n", s)
}
