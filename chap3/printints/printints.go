package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}

func intsToString(values []int) string {
	// 頻繁に文字列を操作する場合はbytesで扱うとメモリ割当の効率がいい
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		// Fprintは型変換を書かなくてもint型を書き込める
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
