package main

import (
	"fmt"
)

func main() {
	s := "💕" // Example string

	// Iterate over each byte in the string
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c",s[i])
	}
	s2 := "❤️" // Example string
  fmt.Println()
	// Iterate over each byte in the string
	for j := 0; j < len(s2); j++ {
		fmt.Println(s2[j])
		//fmt.Printf("%c",s2[i])
	}

}
