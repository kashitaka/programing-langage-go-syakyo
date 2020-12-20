package main

import "fmt"

// この関数は func() int なる型を返す。
// つまり int を返す無名関数が帰る
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
