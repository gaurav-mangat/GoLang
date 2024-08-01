package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello World")
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
	go func() {
		fmt.Println("Hi from the first go routine!!!!!!!")
		wg.Done()
	}()
	time.Sleep(1 * time.Second)

	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())

	go func() {
		fmt.Println("Hi from the second go routine!!!!!!!")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
	fmt.Println("Finally exiting from main goroutine")
	fmt.Println("No. of Goroutines :", runtime.NumGoroutine())
}
