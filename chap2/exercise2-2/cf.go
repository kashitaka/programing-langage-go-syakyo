package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/chap2/tempconv"
)

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			display(t)
		}

	} else {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		t, err := strconv.ParseFloat(stdin.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		display(t)
	}
}

func display(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}
