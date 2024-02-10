package main

type Rectangle struct {
	width  float64
	height float64
}

func Peremeter(rectangle Rectangle) float64 {
	return (rectangle.height + rectangle.width) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.width
}
