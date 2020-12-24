package main

import "fmt"

func main() {
	naturals := make(chan int) // 自然数の送信
	squares := make(chan int)  // n^2 の送信

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		// range を使えば ok が false になるまでloopする
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer メインゴルーチン
	for x := range squares {
		fmt.Println(x)
	}
}
