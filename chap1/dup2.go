// このdupではコマンドライン引数で渡されるファイル名を読む
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// argsがない場合は標準入力をそのまま受け取る
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
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// この関数は呼び出しよりもうしろのブロックにあっても大丈夫
func countLines(f *os.File, counts map[string]int) {
	// 読み込みはinput streamで読むので効率は良い。
	// dup3の例はメモリにファイルを読み込んでいるバージョン
	input := bufio.NewScanner(f)
	for input.Scan() {
		// 行を keyとする数値を加算
		// キーがまだ存在していない場合は初期値0として加算される
		counts[input.Text()]++
	}
}
