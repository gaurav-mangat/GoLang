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
	sl=append(sl,52)
    fmt.Println()
    fmt.Println("Modified Slice  is : ",sl)
    fmt.Println()
	sl=append(sl,53,54,55)
    fmt.Println()
	y:=[] int{56,57,58,59,60}
	sl=append(sl,y...)
    fmt.Println("Modified Slice  is : ",sl)
    fmt.Println()
	
}
