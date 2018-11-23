package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	// 宣言した変数のメモリ領域のアドレスを確保する為の変数
	// var p *int　の省略を使用する
	p := &i // &オペレーターはそのオペランドへのポインタを返す
	fmt.Println(*p)
	fmt.Println(p)
	*p = 21
	fmt.Print("iは")
	fmt.Println(i)
	fmt.Print("pは")
	fmt.Println(*p)

	p = &j
	fmt.Print("jは")
	fmt.Println(j)

	*p = *p / 37
	fmt.Print("jの値")
	fmt.Println(j)
}
