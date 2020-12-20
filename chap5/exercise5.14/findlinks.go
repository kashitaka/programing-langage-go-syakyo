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
	keys := []string{}
	for k := range prereqs {
		keys = append(keys, k)
	}
	breadthFirst(crawl, keys)
}

// 幅優先探索
// worklistに対して f を行う
// f は item に対して何らかの操作を行い、子ノードを返す関数を渡す
// f が worklistに新しい要素を追加し、かつ要素が未探索である限り、
// for len(worklist) > 0 の通り loop し続ける
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(course string) []string {
	fmt.Println(course)
	list := prereqs[course]
	return list
}
