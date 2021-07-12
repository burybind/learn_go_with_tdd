package shapes

import "math"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	baseWidth, height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return r.width*2 + r.height*2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return .5 * t.baseWidth * t.height
}
