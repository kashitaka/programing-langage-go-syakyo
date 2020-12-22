package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type strReader struct {
	s string
}

func (r *strReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	return
}

func NewReader(s string) *strReader {
	r := strReader{s: s}
	return &r
}

// func (r *Reader) Read(b []byte) (n int, err error) {
// 	if r.i >= int64(len(r.s)) {
// 		return 0, io.EOF
// 	}
// 	r.prevRune = -1
// 	n = copy(b, r.s[r.i:])
// 	r.i += int64(n)
// 	return
// }

// 標準入力に渡されたhtml内の全ての a タグのリンクを表示する
func main() {
	r := NewReader("<html><a href=\"aaaa\"></a></html>")

	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
