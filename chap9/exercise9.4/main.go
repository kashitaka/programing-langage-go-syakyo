package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 1000000; {
		i *= 10
		bench(i)
	}
	// 10 routins : 0.000053秒
	// 100 routins : 0.000528秒
	// 1000 routins : 0.004482秒
	// 10000 routins : 0.055743秒
	// 100000 routins : 0.591393秒
	// 1000000 routins : 5.731039秒
	// 10000000 は メモリの利用量が20Gを超えて結果が帰ってこなかった。
}

func bench(concur int) {
	ch := make(chan int)
	start := time.Now()

	go recurs(concur, ch)

	<-ch
	end := time.Now()
	fmt.Printf("%d routins : %f秒\n", concur, (end.Sub(start)).Seconds())
}

func recurs(c int, ch chan int) {
	newCh := make(chan int)
	if c != 0 {
		c--
		go recurs(c, newCh)
		i := <-newCh
		ch <- i
	} else {
		ch <- 9
	}
}
