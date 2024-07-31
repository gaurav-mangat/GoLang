package main

import "fmt"

func main() {
	a := func() (int, int) {
		return 4, 6
	}
	fmt.Println("Addition is : ", add(a))
}
func add(a func() (int, int)) int {
	x, y := a()
	return x + y
}
