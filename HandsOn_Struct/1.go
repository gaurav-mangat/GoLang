package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"chocolate", "banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIC)
	fmt.Println(p2.favIC)

	for _, v := range p1.favIC {
		fmt.Println(p1.first, "favorite is", v)
	}

	for _, v := range p2.favIC {
		fmt.Println(p2.first, "favorite is", v)
	}
}


