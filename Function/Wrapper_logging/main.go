package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintln("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprintln("The count of the book is ", strconv.Itoa(int(c)))
}
func logInfo(s fmt.Stringer) {
	log.Println("Log from me :", s.String())
}
func main() {
	b := book{"History Book"}
	c := count(15)
	fmt.Println(b)
	fmt.Println(c)
	log.Println(b)
	log.Println(c)
	logInfo(b)
	logInfo(c)
}
