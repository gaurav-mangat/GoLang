package main

import "fmt"

func main() {
	var a byte = 'R'
	fmt.Printf("%c is of type %T\n", a, a)
	var b string = "💕"
	fmt.Println(len(b))
	fmt.Printf("%U is of type %T\n", b, b)
	var c byte = '9'
	fmt.Printf("%d is of type %T", c, c)

}
