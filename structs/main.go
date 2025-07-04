package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rectangle Rectangle) Perimeter() float64 {
	width := rectangle.Width
	height := rectangle.Height
	if width < 0 || height < 0 {
		return 0
	}
	return 2 * (width + height)
}

func (rectangle Rectangle) Area() float64 {
	width := rectangle.Width
	height := rectangle.Height
	if width < 0 || height < 0 {
		return 0
	}
	return width * height
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
