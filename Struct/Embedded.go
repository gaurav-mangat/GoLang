package main

import "fmt"

type Student struct {
	name   string
	branch string
	year   int
}

type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

func main() {

	result1 := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{"Bongo", "CSE", 2},
	}
	result2 := Teacher{
		"Aarushi Singhal", "Golang", 3, Student{"Gaurav", "Intern", 1},
	}

	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result1.name)
	fmt.Println("Subject: ", result1.subject)
	fmt.Println("Experience: ", result1.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", result1.details.name)
	fmt.Println("Student's branch name: ", result1.details.branch)
	fmt.Println("Year: ", result1.details.year)

	fmt.Println(result2)
}
