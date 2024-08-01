package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU's\t\t", runtime.NumCPU())
	fmt.Println("Goroutine's\t\t", runtime.NumGoroutine())

	go foo()
	bar()
	fmt.Println("CPU's\t\t", runtime.NumCPU())
	fmt.Println("Goroutine's\t\t", runtime.NumGoroutine())
}
func foo() {
	for i := 1; i < 11; i++ {
		fmt.Println("foo  ", i)
	}
}

func bar() {
	for i := 11; i < 21; i++ {
		fmt.Println("bar  ", i)
	}
}

// Here in the output we can see that bar function is not running because
// before it runs, the main function is ending

//
