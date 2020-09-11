// echo1 はLinuxコマンドechoの実装です
package main

import (
	"fmt"
	"os"
)

func main() {
	// ここでは変数s, sepをstringで定義している。
	// なおstringの初期値は "" である。
	var s, sep string

	// forループの宣言
	// 一般化すると
	// for initialization; condition; post
	// initialization => Option[T] 初期処理
	// condition => bool trueの場合にループが実行される
	// post => Option[T] ループ後の処理
	// この3つは省略可能
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// 無限ループ
	// for {
	// }

	// whileループ(conditionのみを記述)
	// for s != "a" {
	// }

	fmt.Println(s)
}
