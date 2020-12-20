package main

import (
	"fmt"
	"log"
	"sort"
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
	"linear algebra":        {"calculus"},
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
	var visitAll func(items []string, parents []string)
	visitAll = func(items []string, parents []string) {
		for _, item := range items {
			for _, v := range parents {
				if item == v {
					log.Printf("Circulation occurred at: %v -> %s\n", parents, v)
				}
			}
			if !seen[item] {
				seen[item] = true
				visitAll(m[item], append(parents, item))
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys, nil)
	return order
}
