package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	w io.Writer
	c int64
}

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.c += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c, &c.c
}

func main() {
	w, c := CountingWriter(os.Stdout)
	fmt.Fprintln(w, "ABC")
	fmt.Println(*c) // 4 (ABC + \n)

	fmt.Fprintln(w, "XYZ")
	fmt.Println(*c) // 8
}
