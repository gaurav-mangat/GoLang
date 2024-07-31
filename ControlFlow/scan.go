package main

import (
    "fmt"
)

func main() {
    var name string
    var age int

    // Prompt user for input
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name) // Read input from stdin and store it in 'name'

    fmt.Print("Enter your age: ")
    fmt.Scanln(&age) // Read input from stdin and store it in 'age'

    // Print the input received
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

}

