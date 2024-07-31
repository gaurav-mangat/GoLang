package main
import "fmt"


func main(){

p1:=struct{
	first string
	friends map[string] int
	favDrinks []string
} {
	first : "Ritik",
	friends: map[string] int{
		"Gaurav":1,
	"Prince":2,
},
favDrinks: []string{"Tuborg","Kingfisher"},
}
fmt.Println(p1)
}