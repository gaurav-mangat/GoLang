package main

import "fmt"

func main(){
	a:=[] string {"Garurav","Mangat","Intern","Noida"}
	b:=[] string {"Prince","Mangat","Student","Sikar"}
	fmt.Println(a)
	fmt.Println(b)
	aa:=[][] string{a,b}
	ba:=[][] string{{"hi","hello"},{"HI","HELLO"}}
	fmt.Println(aa)
	fmt.Println(ba)
}