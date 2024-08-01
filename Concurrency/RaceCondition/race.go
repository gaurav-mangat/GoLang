package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 10000
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {

			v := counter
			v++
			counter = v
			//time.Sleep(10000000 * 10)

			wg.Done()
		}()

		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
