package main

import "fmt"

type myset interface {
	constraints.Integer | constraints.Float
}

func mul[t myset](a, b t) t {
	return a * b
}

type myalias int

func main() {
	var a myalias = 44
	fmt.Println(mul(a, 6))
	fmt.Println(mul(14, 6.8))
}
