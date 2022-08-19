package main

import "fmt"

func main() {

	//declarations
	var a []int
	primes := []int{2, 3, 5, 7, 11, 13}
	names := []string{}
	fmt.Println(a, primes, names)

	b := make([]int, 10)
	fmt.Printf("slice : %d size: %d  capacity %d \n", b, len(b), cap(b))

	c := make([]int, 10, 20)
	fmt.Printf("slice : %d size: %d  capacity %d \n", c, len(c), cap(c))

	//interations
	numbers := []int{1, 2, 3, 4, 5}

	for i := range numbers {
		numbers[i] *= 2
	}

	fmt.Println(numbers)
}
