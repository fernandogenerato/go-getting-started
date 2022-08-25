package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//sync multiples gouroutine
func main() {

	start := time.Now()
	rand.Seed(start.UnixNano())

	var control sync.WaitGroup

	for i := 0; i < 5; i++ {
		control.Add(1)
		go execute(&control)
	}
	control.Wait()
	fmt.Printf("Terminating in %s. \n", time.Since(start))
}

func execute(s *sync.WaitGroup) {
	defer s.Done()
	elapsed := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Sleeping by %s...\n", elapsed)
	time.Sleep(elapsed)
}
