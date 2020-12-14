package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(rotate(ints, 2))  // [5 6 1 2 3 4]
	fmt.Println(rotate(ints, 3))  // [4 5 6 1 2 3]
	fmt.Println(rotate(ints, 30)) // [1 2 3 4 5 6]
}

func rotate(ints []int, cnt int) []int {
	l := len(ints)
	shift := cnt % l
	new := make([]int, l)
	copy(new[:shift], ints[l-shift:])
	copy(new[shift:], ints[:l-shift])
	return new
}
