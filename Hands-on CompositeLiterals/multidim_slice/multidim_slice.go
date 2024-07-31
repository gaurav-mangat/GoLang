package main

import "fmt"

func main(){
	sll:=[][] string{
		[] string{"James","Bond","Shaken not stirred"},
		[] string{"Miss","Moneypenny","I'm not 008"},
	}
for i,v :=range sll{
	fmt.Println("Slice no. ",i)
	for i,v:=range v{
		fmt.Println(i," - ",v)
	}
	

}
}