package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{
		Height: 10.0,
		Width: 10.0,
	}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func ExamplePerimeter() {
	rect := Rectangle{
		Height: 10.0,
		Width: 10.0,
	}
	perimeter := Perimeter(rect)
	fmt.Println(perimeter)
	// Output: 40
}

// This is called a table driven test, and allows
// for testing multiple cases at once
func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}

func ExampleRectangle() {
	rectangle := Rectangle{12, 6}
	fmt.Println(rectangle.Area())
	// Output: 72
}

func ExampleTriangle() {
	triangle := Triangle{12, 6}
	fmt.Println(triangle.Area())
	// Output: 36
}

func ExampleCircle() {
	circle := Circle{10}
	fmt.Println(circle.Area())
	// Output: 314.1592653589793
}