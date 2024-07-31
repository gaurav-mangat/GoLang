package main

import "fmt"

func divide(a, b float64) float64 {
	return a / b
}
func main() {
	fmt.Printf("The result of division is %f: ", divide(10, 3))
}
