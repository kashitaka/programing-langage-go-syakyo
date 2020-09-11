// echoの改良。stringsパッケージのJoinを使う
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))

	// 体裁を気にしないならこれでもsliceを見れる
	// fmt.Println(os.Args[1:]) [あああ いいい]
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
