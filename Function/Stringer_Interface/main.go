package main


import "fmt"
// import "strconv"

type book struct{
	title string
}

func (b book) String() string{
	return fmt.Sprint("The title of the book is ", b.title)

}
type  count int

func (c count) String() string{
	//return fmt.Sprint("This is the number : ",strconv.Itoa(int(c)))
	return fmt.Sprintf("This is the number : %d ",int(c))
}
func main(){
b:=book{
	title: "West with The Night",
}
var n count = 42
fmt.Println(b)
fmt.Println(n)
}