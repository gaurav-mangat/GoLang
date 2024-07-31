package main

import "fmt"

func foo() int {
	return 68
}

func bar() (int, string) {
	return 111, "Gaurav"
}
func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
