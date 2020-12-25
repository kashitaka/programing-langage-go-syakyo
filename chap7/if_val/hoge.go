package main

import "fmt"

type Hoge interface {
	Fuga() string
}

type hogeImpl string

func (_ hogeImpl) Fuga() string {
	return "fuga"
}

func (_ hogeImpl) Foo() {
	fmt.Println("foo")
}

type dameImpl string

func (_ dameImpl) Fuga() string {
	return "dame"
}

func main() {

	// インターフェース値の宣言
	var a Hoge
	fmt.Println(&a)      // 使わないと怒られるので print
	a = hogeImpl("hoge") // Hoge interfaceを満たす hogeImpl 型は代入可能
	fmt.Println(&a)      // ただコンパイラはまだ a を Hoge インターフェース値として扱う。
	// a.Foo()              なのでこれはコンパイルエラー

	// 型アサーションする
	b := a.(hogeImpl)
	b.Foo() // 呼べる

	// 改めてインターフェース値の宣言
	var p Hoge
	q := p.(hogeImpl) // 型アサーション
	q.Foo()           // 書けるが実行時に上の行でpanic: interface conversion: main.Hoge is nil, not main.hogeImpl

	// 改めて...
	var x Hoge
	x = dameImpl("dame") // これでインターフェース値のvalueはnilではなくなった
	y := x.(hogeImpl)    // もちろん I/F 満たす hogeImpl にアサーションできる
	y.Foo()              // 書けるが実行時に上の行でpanic: interface conversion: main.Hoge is main.dameImpl, not main.hogeImpl
}
