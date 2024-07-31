package main

import "fmt"

// Type aliases
type Celsius float64
type UserID  int

func main() {
    var temperature Celsius = 80
    var userId UserID = 11

    

    fmt.Println("Temperature:", temperature)
    fmt.Println("User ID:", userId)
}
