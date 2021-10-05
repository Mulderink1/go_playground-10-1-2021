package main

import "fmt"

type shape interface{
	getArea() float64
}

type square struct{
	width float64
	height float64
}

type triangle struct{
	base float64
	height float64
}

func main() {
	s := square{
		width: 23.5,
		height: 45.34,
	}

	t := triangle{
		base: 34.54,
		height: 42.33,
	}

	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.width * s.height
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}