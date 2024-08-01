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
	const gs = 1000
	var wg sync.WaitGroup

	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			//time.Sleep(10000000 * 10)
			mu.Unlock()
			wg.Done()

		}()

		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
