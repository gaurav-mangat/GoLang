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
    fmt.Println("Slice is : ",sl)
	sl=append(sl[:3],sl[6:]...)
    fmt.Println("Slice after deletion is : ",sl)
	
    
	
}
