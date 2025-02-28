package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU's\t\t", runtime.NumCPU())
	fmt.Println("Goroutine's\t\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPU's\t\t", runtime.NumCPU())
	fmt.Println("Goroutine's\t\t", runtime.NumGoroutine())
	wg.Wait()
}
func foo() {
	for i := 1; i < 11; i++ {
		fmt.Println("foo  ", i)
	}
	wg.Done()
}

func bar() {
	for i := 11; i < 21; i++ {
		fmt.Println("bar  ", i)
	}
}
