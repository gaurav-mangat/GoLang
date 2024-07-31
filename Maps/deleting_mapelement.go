package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
		"Pat": 24,
		"Tom": 16,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

    delete(am,"Pat")
    delete(am,"Pattt")
	
	fmt.Println(am)
    delete(am,"Vattt")
	fmt.Println(am["Jay"])

}