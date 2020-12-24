package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		val := strings.Split(v, "=")
		if len(val) != 2 {
			log.Fatalf("Invalid arguments: %s", v)
		}
		location := val[0]
		host := val[1]
		conn, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go display(os.Stdout, conn, location)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func display(dst io.Writer, src io.Reader, loc string) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		fmt.Fprintf(dst, "%s :\t%s\n", loc, scanner.Text())
	}
}
