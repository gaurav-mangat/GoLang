package main

import (
	"encoding/json"
	"fmt"
)

type internship struct {
	Name string
	Age  int
	Date string
}

func main() {
	i1 := internship{"Gaurav", 21, "03-07-2-24"}
	i2 := internship{"Anish", 20, "03-07-2-24"}

	i := []internship{i1, i2}

	a, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))
}
