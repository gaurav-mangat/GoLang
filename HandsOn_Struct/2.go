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
m:=map[string] person{
p1.last:p1,
p2.last:p2,
}
fmt.Println(m)

for k,v := range m{
	fmt.Println(k,v)

for _,v2:= range v.favIC{
fmt.Println(v2)
}
}
}
