package main

import (
	"fmt"
	"sync"
	"time"
)

// without mutex time elapsed: 13.3813ms

func main() {
	var counter = 0
	var wg sync.WaitGroup
	wg.Add(50000)
	startTime := time.Now()
	for i := 0; i < 50000; i++ {
		go func() {
			counter = ((counter * 100) + (counter - 100)) % 13
			wg.Done()
		}()

	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println("time elapsed:", endTime.Sub(startTime))
	fmt.Println(counter)
}
