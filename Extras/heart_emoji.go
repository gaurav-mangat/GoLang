package main

import (
	"fmt"
)

func main() {
	s := "💕" // Example string

	fmt.Println("Byte values in hexadecimal:")

	// Iterate over each byte in the string
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune values and their UTF-8 byte sequences:")

	// Iterate over each rune in the string
	for i, runeValue := range s {
		// Print the rune value
		fmt.Printf("Rune: %c (Code Point: %U)\n", runeValue, runeValue)
		// Print the UTF-8 byte sequence of the rune
		fmt.Printf("UTF-8 bytes: ")
		for _, b := range []byte(string(runeValue)) {
			fmt.Printf("%x ", b)
		}
		fmt.Println()
		fmt.Println("Starting byte position:", i)
	}
}
