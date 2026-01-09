package areacalc

import (
	"strings"
)

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	Shape
	side_a float64
	side_b float64
	t      string
}

func (rectangle *Rectangle) Type() string {
	return rectangle.t
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.side_a * rectangle.side_b
}

func NewRectangle(side_a float64, side_b float64, t string) *Rectangle {
	return &Rectangle{side_a: side_a, side_b: side_b, t: t}
}

type Circle struct {
	Shape
	radius float64
	t      string
}

func (circle *Circle) Type() string {
	return circle.t
}

func (circle *Circle) Area() float64 {
	return pi * (circle.radius * circle.radius)
}

func NewCircle(radius float64, t string) *Circle {
	return &Circle{radius: radius, t: t}
}

func AreaCalculator(figures []Shape) (string, float64) {
	var str strings.Builder
	var sum float64
	last_figure := len(figures) - 1

	for i, figure := range figures {
		sum += figure.Area()
		str.WriteString(figure.Type())
		if i != last_figure {
			str.WriteString("-")
		}
	}

	return str.String(), sum
}
