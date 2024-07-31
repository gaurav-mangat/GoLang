package main

import "fmt"

type student struct{
name string
}
func (s student) namee (){
	fmt.Println("Name is : ",s.name)
} 

func main(){
s1:=student{"Gaurav"}
s2:=student{"Alok"}
s1.namee()
s2.namee()
}