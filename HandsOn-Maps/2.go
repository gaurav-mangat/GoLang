package main

import "fmt"
func main(){
	m:=map[string] []string{
		`bond_james`  : {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`  : {`intelligence`, `literature`, `computer science`},
		`no_dr`  : {`cats`, `ice cream`, `sunsets`},
		`fleming_inn`:{`steaks`,`cigars`,`espionage`},
}
for a,b:=range m{
	fmt.Println(a)
	for c,d :=range b{
		fmt.Printf("%v  -   %v\n",c,d)
	}
}
}