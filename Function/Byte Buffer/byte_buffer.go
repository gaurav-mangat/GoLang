package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("HELLO BHAI")
	fmt.Println(b.String())
	b.WriteString(" Gaurav here!!")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("How are you doing?")
	fmt.Println(b.String())

	b.Write([]byte("  Happy Happy"))
	fmt.Println(b.String())
}
