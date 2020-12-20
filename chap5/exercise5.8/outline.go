package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	fmt.Println(ElementById(doc, "page"))
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) != nil {
			return n
		}
	}
	if post != nil {
		if post(n) {
			return n
		}
	}
	return nil
}

var depth int

func ElementById(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}

		}
		return false
	}
	return forEachNode(n, pre, nil)
}
