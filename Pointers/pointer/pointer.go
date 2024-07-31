package main

import "fmt"

func main() {
	a := 70
	b:=7

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(&b)
	
	
	fmt.Printf("%v is of type %T",&a,&a)
}
