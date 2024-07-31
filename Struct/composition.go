package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person  // Embedded struct
    EmployeeID int
    Salary     float64
}

func (p *Person) SayHello() {
    fmt.Println("Hello, my name is", p.Name)
}

func (e *Employee) SayHello() {
	e.EmployeeID+=2
    fmt.Println("Hello, I am an employee with ID after modification: ", e.EmployeeID)
}

func main() {
    emp := Employee{
        Person: Person{
            Name: "John Doe",
            Age:  30,
        },
        EmployeeID: 1234,
        Salary:     50000.0,
    }

    // Accessing embedded fields
    fmt.Println("Employee:", emp.Name) // Output: John Doe
    fmt.Println("Age:", emp.Age)       // Output: 30
    fmt.Println("ID before modification: ",emp.EmployeeID)
    // Calling methods
    emp.SayHello() // Output: Hello, I am an employee with ID 1234
    fmt.Println("ID after  modification in main function : ",emp.EmployeeID)
	
}
