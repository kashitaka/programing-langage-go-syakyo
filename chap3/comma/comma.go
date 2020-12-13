package main

import "fmt"

func main() {
	fmt.Println(comma("1234567890"))  //1,234,567,890
	fmt.Println(comma("Hello world")) //He,llo, wo,rld
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
