package main

import "fmt"

func mul[t int | float64](a, b t) t {
	return a * b
}
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(mul(a, b))
	//fmt.Println(mul(14, 6.8))
}
