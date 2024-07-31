package main

import "fmt"

func main() {
	defer fmt.Println("World") // This will be executed last
	fmt.Println("Hello")

	// Deferred functions are executed in LIFO order
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
	a()
}

	func a() {
    i := 2
    defer fmt.Println(i)
    i++
	fmt.Println(i);
    return
}


