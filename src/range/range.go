package main

import "fmt"

func main() {
	s := []int{2, 3, 8}

	for i, v := range s {
		fmt.Println(i, v)
		// 0 2
		// 1 3
		// 2 8
	}

	for _, v := range s {
		fmt.Println(v)
		// 2
		// 3
		// 8
	}

	m := map[string]int{"Haru": 200, "Miya": 300}

	for k, v := range m {
		fmt.Println(k, v)
		// Haru 200
		// Miya 300
	}
}
