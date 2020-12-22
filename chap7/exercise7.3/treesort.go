package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value int
	// ポインタ型であれば再起でフィールド宣言できる
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		add(root, v)
	}
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	list := appendValues([]int{}, t)
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range list {
		fmt.Fprintf(&buf, "%d", v)
		if i != len(list)-1 {
			fmt.Fprint(&buf, ", ")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	a := tree{value: 5}
	add(&a, 10)
	add(&a, 1)
	add(&a, 100)
	add(&a, 2)

	fmt.Println(a.String())
}
