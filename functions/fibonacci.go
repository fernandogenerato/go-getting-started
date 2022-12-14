package main

import "fmt"

//fib sequence using closures
func main() {
	a, b := 0, 1
	fib := func() int {
		a, b = b, a+b
		return a
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", fib())
	}
}
