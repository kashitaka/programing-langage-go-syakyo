package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("", "a"))                // false
	fmt.Println(isAnagram("", ""))                 // true
	fmt.Println(isAnagram("こんにちは世界", "こんばんは世界"))   // false
	fmt.Println(isAnagram("こんにちは世界🥺", "世界には🥺ちんこ")) // true
}

func isAnagram(s1, s2 string) bool {
	map1 := make(map[rune]int)
	for _, r := range s1 {
		map1[r]++
	}
	map2 := make(map[rune]int)
	for _, r := range s2 {
		map2[r] += 1
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	for k, v := range map2 {
		if map1[k] != v {
			return false
		}
	}
	return true
}
