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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s\n", depth*2, "", n.Data, showAttr(n))
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, showAttr(n))
		}
		depth++
	}
	if n.Type == html.TextNode {
		fmt.Printf("%*s%s", depth*2, "", n.Data)
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func showAttr(n *html.Node) string {
	strAtr := ""
	for _, a := range n.Attr {
		strAtr += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
	}
	return strAtr
}
