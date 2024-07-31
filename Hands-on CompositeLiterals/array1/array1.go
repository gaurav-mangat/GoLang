package main

import "fmt"

func main() {
    arr := [5]int{}
    
    for i := range arr {
        fmt.Println("Enter value for", i, "index of array:")
        fmt.Scan(&arr[i])
    }

    fmt.Println()
    fmt.Println()
    fmt.Println("Array is:")
    for i, v := range arr {
        fmt.Printf("%d th value is %d and is of type %T\n", i, v, v)
    }
}
