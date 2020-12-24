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
		for {
			x, ok := <-naturals
			if !ok {
				// これがないと squareには int の ゼロ値 0 が送られ続ける
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer メインゴルーチン
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
