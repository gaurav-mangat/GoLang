package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v \t\n", x)

	if x <= 100 {
		fmt.Println("X is less than 100\n")
	} else if x > 100 && x <= 200 {
		fmt.Println("Between 101 and 200\n")
	} else if x > 200 && x <= 250 {
		fmt.Println("X is between 201 and 250\n")
	} else {
		fmt.Println("X is more than 250\n")
	}
}
