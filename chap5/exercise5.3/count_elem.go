package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 標準入力に渡されたhtml内の全ての a タグのリンクを表示する
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}
	show(doc)
}

func show(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("%s\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "script" {
			continue
		}
		show(c)
	}
}
