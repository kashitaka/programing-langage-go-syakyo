package main

import (
	"fmt"
)

func main() {
	fmt.Println(KB, MB, GB)
}

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
)
