package main

import (
	"fmt"
	"time"
)

func task1() {
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished!")

}

func task2() {
	fmt.Println("task2 finished!")
}

func onesecond() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("1秒sleep")
}

func twosecond() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2秒sleep")
}

func threesecond() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("3秒sleep")
}

func main() {
	go task1()
	go task2()

	// task2 finished!
	// task1 finished!
	time.Sleep(time.Second * 3)

	fmt.Println("--- goroutins (time確認) ---")
	fmt.Println(time.Now())
	go twosecond() // 並列処理
	go onesecond() // 並列処理
	threesecond()
	fmt.Println(time.Now())
}
