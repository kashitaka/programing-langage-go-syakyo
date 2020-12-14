package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := []byte("aa aaaa     a     aaa")
	fmt.Printf("%s\n", rmAdjDup(str)) // "aa aaaa a aa"
	str = []byte("          ")
	fmt.Printf("%s\n", rmAdjDup(str)) // " "
}

func rmAdjDup(str []byte) []byte {
	idx := 0
	for i := 0; i < len(str)-1; {
		r, size := utf8.DecodeRune(str[i:])
		nextRune, _ := utf8.DecodeRune(str[i+size:])
		if r == nextRune && unicode.IsSpace(r) {
			i += size
			continue
		}
		i += size
		idx += size
		copy(str[idx:idx+size], str[i:i+size])
	}
	return str[:idx]
}
