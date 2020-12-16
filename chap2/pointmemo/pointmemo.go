package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// ポインタ型と & とか * の演算子について動かして確認したのでメモ
	// https://golang.org/ref/spec#Address_operators

	// ポインタ型は通常の型に * をつけることで表せる
	// 例 *int
	var a *int
	fmt.Println(a) // nil

	// & アドレス演算子という。
	// 変数の前につけることでその変数のアドレスを取り出せる
	// the address operation &x generates a pointer of type *T to x.
	// または参照できる別のエイリアスをつけているとも言える
	b := string("aaa")
	c := &b
	fmt.Println(c)

	// * はポインタ間接参照と呼ばれる演算子
	// *T型の変数を示す
	// アドレス変数の参照先からコピーして実体を得るという理解
	d := *c
	d = "bbb"
	fmt.Println(d)

	p := &Person{
		Name: "太郎",
		Age:  20,
	}
	fmt.Printf("p: %v\n", p)
	p2 := *p // アドレスにある変数のコピーをもらう
	p2.Age = 30
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p2: %v\n", p2)
	p3 := p // p3はpそのものの参照。pの別のエイリアス
	p3.Age = 40
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p2: %v\n", p3)

	fmt.Println()

	q := Person{ // pは &Personとすることで無名変数への参照だった。
		Name: "太郎",
		Age:  20,
	}
	fmt.Printf("q: %v\n", q)
	q2 := &q
	q2.Age = 30
	fmt.Printf("q: %v\n", q)
	fmt.Printf("q2: %v\n", q2)

	fmt.Println()

	r := &Person{
		Name: "太郎",
		Age:  20,
	}
	changeAge(*r, 30)
	fmt.Printf("r: %v\n", r)

	changeAge3(*r, 30)
	fmt.Printf("r: %v\n", r)

	changeAge4(r, 30)
	fmt.Printf("r: %v\n", r)

	changeAge2(r, 40)
	fmt.Printf("r: %v\n", r)
}

func changeAge(p Person, age int) *Person {
	// コピーを受け取り、その参照を返す
	p.Age = age
	return &p // "&" は実体から参照を生成する演算子
}

func changeAge2(p *Person, age int) *Person {
	// 参照を受け取り、返す
	p.Age = age
	return p
}

func changeAge3(p Person, age int) Person {
	// コピーを受け取り、返す
	p.Age = age
	return p
}

func changeAge4(p *Person, age int) Person {
	// 参照を受け取り、参照を書き換えて、コピーを返す
	p.Age = age
	return *p // "*" は参照をコピーした実体を得る演算子"
}
