package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&b)

	result := add(a, b)
	fmt.Println("Result is : ", result)
}
func add(x, y int) (sum int) {
	sum = x + y
	return
}
