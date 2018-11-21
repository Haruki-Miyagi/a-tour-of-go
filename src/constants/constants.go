package main

import (
	"fmt"
)

// 定数は := を使って宣言できない
const Pi = 3.14

const (
	Big   = 1 << 100
	small = Big >> 99
)

// 関数はreturnを使用して値を返す
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("--- 定数 ---")
	// 文字(character)、文字列(string)、boolean、数値(numeric)
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println("--- 関数を利用しbit数のチェック ---")
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))  // overflows int
	// -> コンピュターの記憶容量が溢れて使いないという意味
}
