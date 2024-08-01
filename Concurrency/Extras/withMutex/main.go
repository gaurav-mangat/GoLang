package main

import (
	"fmt"
	"sync"
	"time"
)

// with mutex time elapsed: 14.6801ms

func main() {
	var counter = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(50000)
	startTime := time.Now()
	for i := 0; i < 50000; i++ {

		go func() {
			mu.Lock()
			counter = ((counter * 100) + (counter - 100)) % 13
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println("time elapsed:", endTime.Sub(startTime))
	fmt.Println(counter)
}
