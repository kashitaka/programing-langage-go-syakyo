package main

import (
	"fmt"
	"log"
	"os"
	"programing-language-go/chap5/links"
)

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int // worklistへの送信待ちの数

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 送信によりtokenを獲得。仮にtokensのバッファが埋まっていたらここで待たされる
	list, err := links.Extract(url)
	<-tokens // 受信によりtokenを開放
	if err != nil {
		log.Print(err)
	}
	return list
}
