package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10000000; {
		bench(i)
		i *= 10
	}
	// 0.000011秒: 92013.249908 count/sec : 1
	// 0.000013秒: 776759.359950 count/sec : 10
	// 0.000158秒: 631241.399336 count/sec : 100
	// 0.000703秒: 1422135.250751 count/sec : 1000
	// 0.007621秒: 1312088.176524 count/sec : 10000
	// 0.075913秒: 1317293.406767 count/sec : 100000
	// 0.687756秒: 1454004.353434 count/sec : 1000000
	// 7.443764秒: 1343406.374164 count/sec : 10000000

}

func bench(c int) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	start := time.Now()
	go func() {
		for {
			<-ch1
			ch2 <- "ping"
		}
	}()

	ch1 <- "pong"

	for i := 0; i <= c; i++ {
		<-ch2
		ch1 <- "pong"
	}

	end := time.Now()
	fmt.Printf("%f秒: %f count/sec : %d\n", (end.Sub(start)).Seconds(), float64(c)/(end.Sub(start)).Seconds(), c)
}
