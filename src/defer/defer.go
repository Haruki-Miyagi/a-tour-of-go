package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("最初")
	defer fmt.Println("2番目")
	defer fmt.Println("3番目")
	// 出力結果
	// 3番目
	// 2番目
	// 最初
	fmt.Println("遅延 => defer")
}
