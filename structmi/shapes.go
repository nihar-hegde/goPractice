package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func Peremeter(rectangle Rectangle) float64 {
	return (rectangle.height + rectangle.width) * 2
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * math.Pi * c.radius
}
