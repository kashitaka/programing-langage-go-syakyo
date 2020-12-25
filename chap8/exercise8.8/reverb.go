package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var port *int = flag.Int("port", 8000, "tcp port number that server listens")

func main() {
	flag.Parse()
	lister, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	ch := make(chan struct{})
	var receive bool
	go func() {
		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintln(c, "Close because you are silent.")
			c.Close()
		case <-ch:
		}
	}()
	input := bufio.NewScanner(c)
	for input.Scan() {
		if !receive {
			// 最初の1回しか送らない制御。（もっとうまいやり方ありそう
			ch <- struct{}{}
			receive = true
			close(ch)
		}
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
