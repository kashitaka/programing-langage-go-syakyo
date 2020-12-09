package main

import (
	"fmt"
	"time"

	"gopl.io/chap2/exercise23"
	"gopl.io/chap2/exercise24"
	"gopl.io/chap2/exercise25"
	"gopl.io/chap2/popcount"
)

const loopCount = 100000

func main() {
	// 結果は
	// original: 0.000149秒
	// 2.3:  0.006522秒
	// 2.4: 0.016669秒
	// 2.5: 2.918475秒
	measurePopCount()
	measurePopCount2()
	measurePopCount3()
	measurePopCount4()
}

func measurePopCount() {
	start := time.Now()
	for i := uint64(0); i < loopCount; i++ {
		popcount.PopCount(i)
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func measurePopCount2() {
	start := time.Now()
	for i := uint64(0); i < loopCount; i++ {
		exercise23.PopCount(i)
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func measurePopCount3() {
	start := time.Now()
	for i := uint64(0); i < loopCount; i++ {
		exercise24.PopCount(i)
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

func measurePopCount4() {
	start := time.Now()
	for i := uint64(0); i < loopCount; i++ {
		exercise25.PopCount(i)
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
