package main

import (
	"fmt"
	"math"
)

type rectange struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s rectange) area() float64 {
	return s.length * s.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("The area of the shape is : ", s.area())
}
func main() {
	a := rectange{length: 5, width: 10}
	b := circle{radius: 5}
	info(a)
	info(b)
}
