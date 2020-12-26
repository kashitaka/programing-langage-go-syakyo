package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// 送信用チャネルの型
type client struct {
	who     string
	channel chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	// 接続クライアントの管理map
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.channel <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var buf bytes.Buffer
			buf.WriteString("Current member: ")
			for c := range clients {
				buf.WriteString(c.who + " ")
			}
			cli.channel <- buf.String()
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	// 他の人のメッセージの待ち受けサブルーチン
	go clientWriter(conn, ch)
	ch <- "Input your name before entering a chatroom"
	input := bufio.NewScanner(conn)
	var who string
	for input.Scan() {
		who = input.Text()
		break
	}

	ch <- "You are " + who
	cl := client{who, ch}
	messages <- who + " has arrived"
	entering <- cl

	resetTimer := make(chan struct{})
	go closeTimer(conn, ch, resetTimer)
	for input.Scan() {
		resetTimer <- struct{}{}
		messages <- who + ": " + input.Text()
		go closeTimer(conn, ch, resetTimer)
	}

	leaving <- cl
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func closeTimer(conn net.Conn, ch <-chan string, rst chan struct{}) {
	select {
	case <-time.After(5 * time.Minute):
		conn.Close()
	case <-rst:
		// do nothing
	}
}
