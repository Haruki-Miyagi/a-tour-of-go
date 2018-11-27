package main

import "fmt"

func main() {
	m := map[string]int{"Haru": 100, "Miya": 200} // map[Haru:100 Miya:200]

	v, ok := m["Haru"]

	fmt.Println(v)  // 100
	fmt.Println(ok) // true

	delete(m, "Haru")
	fmt.Println(m) // map[Haru:200]
}
