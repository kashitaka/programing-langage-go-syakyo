package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890")) //1,234,567,890
}

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	// single byte characterしかこない想定
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if i%3 == 0 && i < n-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
