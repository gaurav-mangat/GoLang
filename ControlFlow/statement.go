package main

import "fmt"

func main() {
	y := 9
	z := -78
	if x := 5; x < y {
		fmt.Println(x)
	} else if x > z {
		fmt.Println(z)
	} else {
		fmt.Println(y)
	}
}
