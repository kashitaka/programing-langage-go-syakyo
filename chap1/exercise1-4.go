// このdupではコマンドライン引数で渡されるファイル名を読む
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// argsがない場合は標準入力をそのまま受け取る
		// この時の *os.File.Name() は /dev/stdin
		countLines(os.Stdin, counts)
	} else {
		// argsを読み取って、指定されたファイル全てを調べる
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, nameAndCount := range counts {
		count := 0
		chainedFileName := ""
		for name, c := range nameAndCount {
			chainedFileName += name + " "
			count += c
		}
		if count > 1 {
			fmt.Printf("%d\t%s\t%s\n", count, line, chainedFileName)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]
		}
		counts[input.Text()][f.Name()]++
	}
}
