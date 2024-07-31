package main

import "fmt"

func main(){
	m:=map[string] int{
		"Gaurav" : 21,
		"Prince" : 18,
		"Ritik" : 21,
}
	fmt.Println("The age of Gaurav is ", m["Gaurav"])

	fmt.Println("-------------------------------------------------")


    m2 := make(map[string]int)
    m2["apple"] = 5
    m2["banana"] = 3
    m2["cherry"] = 8
	m2["Orange"]=15
    fmt.Println("m2:", m2)

    // Accessing individual elements
    fmt.Println("Value of banana:", m2["banana"])
	fmt.Printf("%#v",m2)
}

