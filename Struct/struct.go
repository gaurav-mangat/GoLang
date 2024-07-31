package main

import "fmt"

type student struct {
	name    string
	rollno  int
	class   int
	section string
}

func main() {
	p1 := student{
		name:    "Gaurav",
		rollno:  44,
		class:   12,
		section: "A",
	}
	p2 := student{
		name:    "Ritik",
		rollno:  56,
		class:   12,
		section: "B",
	}
	p2.class = 10
	fmt.Printf("%#v is of type %T\n", p1, p1)
	fmt.Print(p2)
}
