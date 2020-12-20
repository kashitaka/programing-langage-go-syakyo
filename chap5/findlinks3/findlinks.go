package main

import (
	"fmt"
	"log"
	"os"
	"programing-language-go/chap5/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
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

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
