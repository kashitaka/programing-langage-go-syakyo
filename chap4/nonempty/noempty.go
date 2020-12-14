package main

import "fmt"

func main() {
	sl := []string{"a", "b", "c", "", "d"}
	fmt.Println(nonempty(sl)) // [a b c d]
	fmt.Println(sl)           // [a b c d d] nonemptyは基底配列を書き換える
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
