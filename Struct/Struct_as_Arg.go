package main
import "fmt"

type Person struct{
	name string
	age int 
	job string
	salary float32
}

func main(){
	var a Person
	a=Person{
		name:"Gaurav",
		age: 22,
		job:"Intern",
		salary:30000,
	}

	fmt.Printf("%#v is of type %T",a,a)
	fmt.Printf("\nRecord of the employees : ")
	printPerson(a)

	fmt.Print("New name is : ",a.name)
	
}
func printPerson(n Person){
	n.name="Prince"
	fmt.Println("Name :",n.name)
	fmt.Println("Age :",n.age)
	fmt.Println("Job :",n.job)
	fmt.Println("Salary :",n.salary)
}