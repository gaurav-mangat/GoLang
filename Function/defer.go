package main

import "fmt"

func main() {
    fmt.Println("Start")

    // Deferred function call
    defer fmt.Println("Deferred call")

    fmt.Println("End")
}
