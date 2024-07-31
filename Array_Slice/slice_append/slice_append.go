package main

import "fmt"

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("-------------")
	// variadic parameter
	xi = append(xi, 45, 46, 777)
	fmt.Println(xi)
	fmt.Println("-------------")
}
