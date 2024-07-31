package main

import (
    "fmt"
    "math"
)

// Define a geometry interface
type geometry interface {
    area() float64
    perimeter() float64
}

// Define a struct for a rectangle
type rectangle struct {
    width, height float64
}

// Define methods for the rectangle struct to implement the geometry interface
func (r rectangle) area() float64 {
    return r.width * r.height
}

func (r rectangle) perimeter() float64 {
    return 2*r.width + 2*r.height
}

// Define a struct for a circle
type circle struct {
    radius float64
}

// Define methods for the circle struct to implement the geometry interface
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

// Function that takes a geometry interface as an argument
func measure(g geometry) {
    fmt.Println("Shape:", g)
    fmt.Println("Area:", g.area())
    fmt.Println("Perimeter:", g.perimeter())
}

func main() {
    // Create instances of rectangle and circle
    r := rectangle{width: 3, height: 4}
    c := circle{radius: 5}

    // Call measure function with rectangle and circle instances
    measure(r)
    measure(c)
}
