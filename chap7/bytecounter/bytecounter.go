package main

import "fmt"

type ByteCounter int

// io.Writer の interface である Write(p []byte) (n int, err error) メソッドを実装する
// Javaのように明示的に interface を追加するというより、
// io.Writerの仕様書を読んで開発者が同じメソッドを実装するイメージ
func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	name := "Dolly"
	// ByteCounter は io.Writerの I/F を実装しているので
	// Fprintf に入れることができる
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12 = len("hello, Dolly")

	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 24
}
