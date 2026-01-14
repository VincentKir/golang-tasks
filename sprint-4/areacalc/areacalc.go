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
	sideA float64
	sideB float64
	t     string
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.sideA * rectangle.sideB
}

func (rectangle Rectangle) Type() string {
	return rectangle.t
}

var _ Shape = (*Rectangle)(nil)

func NewRectangle(sideA float64, sideB float64, t string) *Rectangle {
	return &Rectangle{sideA: sideA, sideB: sideB, t: t}
}

type Circle struct {
	radius float64
	t      string
}

func (circle Circle) Type() string {
	return circle.t
}

func (circle Circle) Area() float64 {
	return pi * (circle.radius * circle.radius)
}

var _ Shape = (*Circle)(nil)

func NewCircle(radius float64, t string) *Circle {
	return &Circle{radius: radius, t: t}
}

func AreaCalculator(figures []Shape) (string, float64) {
	str := make([]string, 0)
	var sum float64

	for _, figure := range figures {
		sum += figure.Area()
		str = append(str, figure.Type())
	}

	return strings.Join(str, "-"), sum
}
