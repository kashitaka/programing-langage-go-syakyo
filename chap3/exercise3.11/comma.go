package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma_float("-1234567.89"))    // -1,234,567.89
	fmt.Println(comma_float("1234567.890123")) // 1,234,567.890123
	fmt.Println(comma_float("-1.89"))          //-1.89
	fmt.Println(comma_float("-1"))             // -1
	fmt.Println(comma_float("0"))              // 0
}

func comma_float(s string) string {
	dot := strings.LastIndex(s, ".")
	minus := strings.LastIndex(s, "-")
	if dot < 0 {
		// 整数の場合の処理
		return s[:minus+1] + comma(s[minus+1:])
	}
	return s[:minus+1] + comma(s[minus+1:dot]) + s[dot:]
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
