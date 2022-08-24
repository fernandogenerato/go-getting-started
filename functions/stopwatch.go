package main

import (
	"fmt"
	"time"
)

//high order function example using fibonacci sequence
func main() {
	cron(generateFibonacci(10))
	cron(generateFibonacci(100))
	cron(generateFibonacci(50))
}

func cron(f func()) {
	start := time.Now()
	f()
	fmt.Printf("\nElapsed Time : %s\n\n", time.Since(start))
}

func generateFibonacci(n int) func() {
	return func() {
		a, b := 0, 1
		fib := func() int {
			a, b = b, a+b
			return a
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}
