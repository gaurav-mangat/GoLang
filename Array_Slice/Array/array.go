package main

import "fmt"

func main() {
	var arr [4]int
	var arr4 [4]int
	
	arr[2] = 65.0
	fmt.Println(arr)
	fmt.Println(arr4)

	arr2 := [...]string{"Hi", "My", "Name", "is", "Gaurav"}
	fmt.Println(arr2)
	fmt.Println(len((arr2)))

}
