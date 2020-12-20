package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 標準入力に渡されたhtml内の全ての a タグのリンクを表示する
func main() {
	res := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}
	count(res, doc)
	for k, v := range res {
		fmt.Printf("%s\t%d\t\n", k, v)
	}
}

func count(res map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		res[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(res, c)
	}
}
