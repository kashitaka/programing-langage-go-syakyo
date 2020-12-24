package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006/01/02 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
