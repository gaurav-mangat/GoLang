package main

import "fmt"

func main() {
    sl := make([]int,10)
    a:=42
    for i := range 10 {
      sl[i]=a
      a++  
    }

	sl1:= sl[:5]
	sl2:= sl[5:]
	sl3:= sl[2:7]
	sl4:= sl[1:6]
    fmt.Println()
    fmt.Println("Slice 1 is : ",sl1)
    fmt.Println()
    fmt.Println("Slice 2 is : ",sl2)
    fmt.Println()
    fmt.Println("Slice 3 is : ",sl3)
    fmt.Println()
    fmt.Println("Slice 4 is : ",sl4)
    fmt.Println()
	
}
