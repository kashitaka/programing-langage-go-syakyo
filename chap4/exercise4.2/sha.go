package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bit = flag.Int("b", 256, "bits")

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	flag.Parse()
	display(stdin.Text())
}

func display(t string) {
	if *bit == 256 {
		fmt.Fprintf(os.Stdout, "%x\n", sha256.Sum256([]byte(t)))
	} else if *bit == 384 {
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum384([]byte(t)))
	} else if *bit == 512 {
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum512([]byte(t)))
	} else {
		fmt.Fprintln(os.Stderr, "bit argument must be 256, 384 or 512")
	}
}
