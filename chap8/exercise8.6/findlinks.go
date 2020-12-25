package main

import (
	"flag"
	"fmt"
	"log"
	"programing-language-go/chap5/links"
)

type work struct {
	url   string
	depth int
}

var depth *int = flag.Int("d", 0, "crawl depth")
var url *string = flag.String("u", "", "crawl target url")

func main() {
	flag.Parse()
	worklist := make(chan []work)
	unseenLinks := make(chan work)
	go func() { worklist <- []work{work{*url, 1}} }()

	for i := 0; i < 20; i++ {
		go func() {
			for w := range unseenLinks {
				foundLinks := crawl(w.url)
				newworks := []work{}
				for _, v := range foundLinks {
					newworks = append(newworks, work{v, w.depth + 1})
				}
				go func() { worklist <- newworks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, w := range list {
			if *depth > 0 && w.depth > *depth {
				continue
			}
			if !seen[w.url] {
				seen[w.url] = true
				unseenLinks <- w
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
