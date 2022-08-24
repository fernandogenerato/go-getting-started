package main

import "fmt"

func main() {

	i, p := make(chan int), make(chan int)
	ready := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}
	go separate(nums, i, p, ready)

	var odd, pairs []int
	end := false

	for !end {
		select {
		case n := <-i:
			odd = append(odd, n)

		case n := <-p:
			pairs = append(pairs, n)

		case end = <-ready:
		}

	}
	fmt.Printf("odd : %v | Pairs : %v\n", odd, pairs)
}

func separate(nums []int, i, p chan<- int, ready chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	ready <- true

}
