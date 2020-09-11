// dup: 重複行のカウントを表示する
// cat ファイル名 | go run dup1.go
// のようにパイプで入れよう
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// input.Scan()は実行するとカーソルをずらす
	// 読み込むものがあればtrueを返す。
	// 読み込む内容はinput.Text()で取れる
	for input.Scan() {
		// 行を keyとする数値を加算
		// キーがまだ存在していない場合は初期値0として加算される
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
