package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(CountDiff(c1, c2))
}

func CountDiff(c1, c2 [32]byte) uint8 {
	c := uint8(0)
	for i, b := range c1 {
		c += PopCount(byte(b) ^ byte(c2[i]))
	}
	return c
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint8) uint8 {
	return uint8(pc[byte(x>>(0))])
}
