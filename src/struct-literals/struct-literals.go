package main

import "fmt"

type Vertex struct {
	X, Y int
}

type user struct {
	name  string
	score int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p, *p)

	// newでuser構造体分の領域を確保して、初期化して、そのポインタを返す
	u1 := new(user)
	u1.name = "Haru"
	u1.score = 100
	fmt.Println(u1)
	fmt.Println(*u1)
}
