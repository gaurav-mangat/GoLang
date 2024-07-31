package main

import "fmt"

func main() {
    sl := make([]int,10)
    a:=42
    for i := range 10 {
      sl[i]=a
      a++  
    }

    fmt.Println()
    fmt.Println()
    fmt.Println("Slice is:\n")
    for i, v := range sl {
        fmt.Printf("%d th value is %d and is of type %T\n", i, v, v)
    }
}
