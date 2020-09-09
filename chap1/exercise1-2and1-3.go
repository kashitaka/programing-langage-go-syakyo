package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for idx, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(idx, arg)
	}
}

// 1-3
// for  4.133e-05
// Join 3.875e-05
