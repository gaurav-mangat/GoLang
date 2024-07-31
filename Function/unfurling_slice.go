package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("The sum is", sum(xi...))
}


func sum(nn ...int) int {
	fmt.Println(nn)
	fmt.Printf("%T\n", nn)

	n := 0
	for _, v := range nn {
		n += v
	}
	return n
}
