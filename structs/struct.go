package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func(r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func(t Triangle) Area() float64 {
	return t.Width * t.Height * 0.5
}

func(c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}