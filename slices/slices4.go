package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubles := make([]int, len(numbers))

	copy(doubles, numbers)

	fmt.Println(numbers)

	for i := range doubles {
		doubles[i] *= 2
	}
	fmt.Println(doubles)

}
