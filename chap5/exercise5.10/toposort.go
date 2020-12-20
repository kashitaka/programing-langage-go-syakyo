package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithm": {"data structures"},
	"calculus":  {"linear algebra"},
	"compilers": {
		"data structures",
		"formal language",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal language":       {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// まず宣言しないと再起で呼べない。visitAll内から見つけられない
	var visitAll func(item []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				// fmt.Println(item)

				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for k := range m {
		visitAll([]string{k})
	}
	return order
}
