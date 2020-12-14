package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := []byte("abcde")
	fmt.Printf("%s\n", strRev(str)) // "edcba"
	str = []byte("あ1n🥚a")
	fmt.Printf("%s\n", strRev(str)) // "a🥚n1あ"
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func strRev(str []byte) []byte {
	// 文字単位でbyteを逆順にして、最後に全体を逆順
	for i := 0; i < len(str); {
		_, size := utf8.DecodeRune(str[i:])
		reverse(str[i : i+size])
		i += size
	}
	reverse(str)
	return str
}
