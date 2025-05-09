package structures

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	// Arrange
	rectangle := Rectangle{10.0, 10.0}
	expect := 40.0

	// Act
	result := Perimeter(rectangle)

	// Assert
	if result != expect {
		t.Errorf("result %.2f expect %.2f", result, expect)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expect: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expect: 36.00},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expect {
				t.Errorf("result %.2f expect %.2f", result, tt.expect)
			}
		})
	}
}
