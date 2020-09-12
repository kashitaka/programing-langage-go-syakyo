package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 引数の分だけサブルーチンを開始
	}
	for range os.Args[1:] {
		// chチャンネルからの受信を試みる
		// fetch関数内でchへの送信がなされるまで待たされる。
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	fmt.Println("start", url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	f, _ := os.OpenFile("./hoge.txt", os.O_RDWR|os.O_CREATE, 0664)
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", secs, nbytes, url) // chチャンネルへ送信
}
