package main

import "fmt"

func main() {
	var name string
	name = "Gaurav Chaudhary"

	a, _ := fmt.Print(&name)
	fmt.Println(a)
}
