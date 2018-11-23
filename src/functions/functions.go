package main

import "fmt"

// --- point ---
// 返り値の型は必須
// 引数の型は必須
func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return x, y
}

func add_square(x, y int) int {
	x = x + y // 4
	s := 2
	// return は必要
	return square(x, s)
}

func square(x, s int) int {
	return x * s // 8
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(add_square(1, 3))
}
