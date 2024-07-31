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
	str := `[{"First":"Prince","Last":"Mangat","Age":18}]`

	ss := []byte(str)
	var a []person
	var ab []person

	err = json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println("__________________________________________________")
	err = json.Unmarshal(ss, &ab)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ab)

}
