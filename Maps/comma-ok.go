package main 
import "fmt"

func main(){
	m:=map[string] int{
		"A":1,
		"B":2,
		"C":3,
		"D":4,
	}

	// v,ok:=m["C"]
	// if ok{
	// 	fmt.Println("Key-Value pair exist and value is ",v)
	// }else {
	// 	fmt.Println("Not exist.....")
	// }


	if v, ok:= m["B"]; !ok {
		fmt.Println("Not exist.....")
		}else {
			
		fmt.Println("Key-Value pair exist and value is ",v)
	}

}