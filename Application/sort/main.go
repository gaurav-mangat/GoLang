package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 2, 9, 1, 5, 6}
	sort.Ints(nums)
	fmt.Println("Sorted integers:", nums)
}
