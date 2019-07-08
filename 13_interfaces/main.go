package main

import (
	"fmt"
	"math"
)

// Shape ...
// Define interface
type Shape interface {
	area() float64
}

// Circle ...
type Circle struct {
	x, y, radius float64
}

// Rectangle ...
type Rectangle struct {
	width, height float64
}

// Value receiver. Only calculating, not changing anything.
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Reactangle Area: %f\n", getArea(rectangle))
}
