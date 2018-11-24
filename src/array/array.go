package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Array ---")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 初期値の設定
	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes) // 1,2,3,4,5,6

	fmt.Println("--- slices ---")
	fmt.Println(primes[1:4]) //2,3,4
	fmt.Println(primes[1:])  //2, 3, 4, 5, 6
	fmt.Println(primes[:4])  //1, 2, 3, 4
	fmt.Println(primes[:])   // 1,2,3,4,5,6
}
