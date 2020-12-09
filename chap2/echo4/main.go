package main

import (
	"flag"
	"fmt"
	"strings"
	"tempconv"
)

// flag.Bool returns *bool type
var n = flag.Bool("n", false, "omit training newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
