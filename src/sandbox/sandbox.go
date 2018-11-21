package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcom to the playground!")
	fmt.Println("The time is", time.Now())

	t, _ := time.ParseDuration("1h30m")
	fmt.Println(t.Minutes()) // 90
}
