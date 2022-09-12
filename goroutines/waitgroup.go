package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
  
	wg.Add(3)
  
	routines("A", 3, &wg)
	routines("C", 2, &wg)
	routines("B", 1, &wg)
  
	wg.Wait()
  
}

func routines(action string, sleep time.Duration, wg *sync.WaitGroup) {
  
	go func() {
		start := time.Now()
		for i := 0; i < 3; i++ {
			time.Sleep(sleep * time.Second)
		}
		fmt.Println("action : " + action)
		fmt.Println("time elapsed : ", time.Since(start))
    
		wg.Done()
	}()
}
