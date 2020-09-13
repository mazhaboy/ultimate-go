package main

import (
	"fmt"
)

type Geometry interface {
	Area() float64
	Perimetr() float64
}

type Circle struct {
	r float64
}
type Rectangle struct {
	a float64
	b float64
}

func (r *Circle) Area() float64 {
	return 3.14 * r.r * r.r
}
func (r *Rectangle) Area() float64 {
	return r.a * r.b
}

func main() {
	circ := Circle{
		r: 5,
	}
	d := Rectangle{
		a: 5,
		b: 6,
	}
	fmt.Println(circ.Area())
	fmt.Println(d.Area())

}
