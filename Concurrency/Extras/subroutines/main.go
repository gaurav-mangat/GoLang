package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

func main() {
	var wg sync.WaitGroup
	const os = 500
	const ir = 50

	var mu sync.Mutex
	starttime := time.Now()
	for i := 0; i < os; i++ {
		wg.Add(1)
		go func() {
			var innerwg sync.WaitGroup
			for j := 0; j < ir; j++ {

				innerwg.Add(1)
				go func() {
					mu.Lock()
					v := counter
					v++
					counter = v
					mu.Unlock()
					//fmt.Println("Coutner :", counter)
					//fmt.Println("Goroutines: ", runtime.NumGoroutine())
					innerwg.Done()
				}()

				innerwg.Wait()
			}
			wg.Done()
		}()

	}

	wg.Wait()
	endtime := time.Since(starttime)

	fmt.Println("Counter:", counter)
	fmt.Println("Now the program is going to end")
	fmt.Println("Time elapsed:", endtime)
}
