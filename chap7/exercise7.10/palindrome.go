package main

import (
	"fmt"
	"sort"
)

type mySort struct {
	ints []int
}

func (s *mySort) Len() int           { return len(s.ints) }
func (s *mySort) Swap(i, j int)      { s.ints[i], s.ints[j] = s.ints[j], s.ints[i] }
func (s *mySort) Less(i, j int) bool { return s.ints[i] < s.ints[j] }

func isParindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if !(!s.Less(i, s.Len()-1-i) && !s.Less(s.Len()-i-1, i)) {
			return false
		}
	}
	return true
}

func main() {
	ints1 := []int{1, 2, 3, 3, 2, 1}
	fmt.Println(isParindrome(&mySort{ints: ints1}))

	ints2 := []int{1, 2, 3, 2, 1}
	fmt.Println(isParindrome(&mySort{ints: ints2}))

	ints3 := []int{1, 2, 3, 4, 5}
	fmt.Println(isParindrome(&mySort{ints: ints3}))

}
