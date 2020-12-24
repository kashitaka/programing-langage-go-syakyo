package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		// ctrl + D　送るまでゴルーチンは ↓ の行で待たされる。
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	// ctrl + D　送るまでメインゴルーチンは ↓ の行で待たされる。
	mustCopy(conn, os.Stdin)
	conn.Close()
	// channelに送信があるまで次の行で待機する。
	<-done
	// done がない場合は go func 内の Println が実行される前に
	// main が終わる可能性がある。（逆にPrintlnが実行される可能性もある)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
