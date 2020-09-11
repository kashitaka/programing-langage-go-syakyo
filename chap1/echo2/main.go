// echo1の別の表現
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	// forの別の記法
	// rangeが二つの変数を返す。
	// 1つ目にindex, 2つ目に値
	// indexはこのループでは利用しない。
	// 利用しない変数はtmpなどとしてもいいが、
	// Goは使わないローカル変数を許さないので _ としないとコンパイルエラー
	// _ はブランク識別子という
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
