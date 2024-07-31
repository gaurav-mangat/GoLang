package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main(){
	p1 := struct {
		first string
		last  string
		age   string
	}{
		first: "James",
		last:  "Bond",
		age:   "Fifteen",
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Printf("%v is of type %T \n", p1,p1)
	fmt.Printf("%v is of type %T \n", p2,p2)
}
