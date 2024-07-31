package main

import "fmt"

func outer() func() int {
	return func() int {
		return 42
	}
}

func main() {
	f := outer()
	fmt.Println(f())
}
