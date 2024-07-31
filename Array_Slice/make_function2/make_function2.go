package main

import "fmt"

func main() {
    // Creating a slice of integers with length 3 and capacity 5
    slice := make([]int,3, 5)

    // Adding elements to the slice
    slice[0] = 1
    slice[1] = 2
    slice[2] = 3

    // Print the slice and its length and capacity
    fmt.Println("Slice:", slice)
    fmt.Println("Length:", len(slice))     // Output: Length: 3
    fmt.Println("Capacity:", cap(slice))   // Output: Capacity: 5
}
