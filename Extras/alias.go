package main

import "fmt"

// Type aliases
type Celsius = float64
type UserID = int

func main() {
    var temperature Celsius = 22.5
    var userId UserID = 1001

    // Type aliases can be used interchangeably with their base types
    var temp float64 = temperature
    var uid int = userId

    fmt.Println("Temperature:", temperature)
    fmt.Println("User ID:", userId)
    fmt.Println("Temperature:", temp)
    fmt.Println("User ID:", uid)
}
