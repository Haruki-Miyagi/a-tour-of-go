package main

import "fmt"

func main() {
	score := 61

	if score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("soso")
	}

	if score := 52; score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("soso")
	}
}
