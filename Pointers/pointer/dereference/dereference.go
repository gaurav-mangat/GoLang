package main

import "fmt"

func main() {
	a := 70

	fmt.Println(a)
	fmt.Println(&a)

	b:=&a
	fmt.Println(b)
	fmt.Println(&b)
	
	
	fmt.Printf("%v is of type %T\n",&a,&a)
	fmt.Printf("%v is of type %T\n",b,b)
	fmt.Printf("%v is of type %T\n",&b,&b)

	fmt.Println(*b )
	*b=1111
	fmt.Println(*b )
}
