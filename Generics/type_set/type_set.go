package main

import "fmt"

type myset interface {
	int | float64
}

func mul[t myset](a, b t) t {
	return a * b
}
func main() {
	fmt.Println(mul(4, 6))
	fmt.Println(mul(14, 6.8))
}
