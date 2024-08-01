package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Num of CPUs:", runtime.NumCPU())
	fmt.Println("Num of Goroutines:", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			//time.Sleep(10000000 * 10)
			wg.Done()

		}()

		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
