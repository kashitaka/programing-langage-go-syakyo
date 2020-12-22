package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var w WordCounter
	fmt.Fprint(&w, "hello world")
	fmt.Println(w) // 2
	fmt.Fprint(&w, "I AM A HERO")
	fmt.Println(w) // 6 = 2 + 4

	var l LineCounter
	text := "こんにちは\nこんばんわ"
	fmt.Fprint(&l, text)
	fmt.Println(l) // 2
	text = "昨日\n今日\n明日"
	fmt.Fprint(&l, text)
	fmt.Println(l) // 5 = 2 + 3
}
