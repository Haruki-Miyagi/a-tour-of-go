package main

import (
	"fmt"
)

func foo(m string) func(string) {
	s := m

	f := func(n string) {
		fmt.Println(s + n)
	}
	return f
}

func main() {
	hello := foo("hello ")
	bye := foo("bye ")
	hello("Taro") //-> hello Taro
	bye("Hanako") //-> bye Hanako
	hello("Jiro") //-> hello Jiro
}
