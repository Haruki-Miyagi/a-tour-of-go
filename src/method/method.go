package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
}

func (u user) show() {
	fmt.Printf("name: %s, socre: %d\n", u.name, u.score)
}

func (u *user) point() {
	u.score++
}

func main() {
	u := user{name: "Haru", score: 100}
	u.show()  //name: Haru, socre: 100
	u.point() // u.score = u.score + 1
	u.show()  // name: Haru, socre: 101
}
