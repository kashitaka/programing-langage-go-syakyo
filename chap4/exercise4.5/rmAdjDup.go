package main

import "fmt"

func main() {
	fmt.Println(rmAdjDup([]string{"a", "a", "b", "c", "d"}))
	fmt.Println(rmAdjDup([]string{"a", "a", "a", "a", "b", "c", "d"}))
}

func rmAdjDup(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s == strings[i] {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}
