package main

import "fmt"

func main() {
	fmt.Println(foo())
}

func foo() (i int) {

	defer func() {
		recover()
		i = 10
	}()

	panic(10)
}
