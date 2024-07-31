package main

import "fmt"

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}


	fmt.Printf("%v is of type %T \n",p1, p1)

}
