package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w1 = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	// #v アドヴァーブ
	fmt.Printf("%#v\n", w1) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
	fmt.Printf("%v\n", w1)  // {{{8 8} 5} 20}

	w1.X = 10
	fmt.Printf("%#v\n", w1)
}
