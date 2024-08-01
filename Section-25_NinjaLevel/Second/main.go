package main

import (
	"encoding/json"
	"fmt"
)

// Define a concrete struct that matches the structure of your JSON data.
type Person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	// Convert the JSON string to a byte slice.
	a := []byte(s)

	// Declare a slice of the concrete struct type.
	var people []Person

	// Unmarshal the JSON into the slice of Person structs.
	err := json.Unmarshal(a, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the result.
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("Person ", i, ":", v.First, v.Last, v.Age)
		for _, b := range v.Sayings {
			fmt.Println("\t", b)
		}
		fmt.Println()
		fmt.Println()
	}
}
