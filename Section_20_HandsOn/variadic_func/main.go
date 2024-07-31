package main

import "fmt"

func main() {
	xi := []int{2, 5, 3, 3, 77, 3, -33, 78}
	fmt.Println("Sum of the elements of the slice is : ", foo(xi...))
	fmt.Println("Sum of the elements of the slice is : ", bar(xi))
}
func foo(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func bar(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
