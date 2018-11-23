package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs((z*z)-x) > 0.01 { // 0.01より小さくすれば 適切な値に近づく(math.Sqrt(2))
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println("--- for(基本型) ---")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("--- for(セミコロンなし & 省略3) ---")
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println("--- for(無限ループ処理) ---")
	// for {
	// 	fmt.Println("Hi!!!")
	// }

	fmt.Println("--- 2の平方根 ---")
	fmt.Println(Sqrt(2))      // A -> 1.4166666666666667
	fmt.Println(math.Sqrt(2)) // A -> 1.4142135623730951
}
