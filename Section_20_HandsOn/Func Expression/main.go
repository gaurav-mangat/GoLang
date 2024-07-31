package main

import "fmt"

func main() {
	s := a("Gaurav Singh Mangat")
	fmt.Println(s)

	y := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i + 1)
		}
	}
	y()
}

var a = func(s string) string {
	return fmt.Sprintln("My name is ", s)
}
