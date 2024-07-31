package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		"James", "Bond", 25,
	}
	p2 := person{
		"James", "Bond", 25,
	}
	p := []person{p1, p2}
	fmt.Println(p)
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	//fmt.Printf("%c", b)

}
