package main

import "fmt"
//import "strconv"

func main(){
	 s  :="Gaurav Singh"
	fmt.Println(s)
	// v,err:=strconv.Atoi(s)
	// if err!=nil{
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(v)
	// }
	for _,v:=range s{
		fmt.Printf("%c is of type%T\n",v,v)
	}
	fmt.Printf("%c",s[2])

}