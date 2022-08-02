/*
//sort integers numbers

$ go run quicksort.go 99 2 31 33 10 5 100 6 1
[1 2 5 6 10 31 33 99 100]
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s not is a valid number!\n", n)
			os.Exit(1)
		}
		numbers[i] = num
	}
	fmt.Println(quicksort(numbers))

}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)

	middleIndex := len(sortedNumbers) / 2
	index := sortedNumbers[middleIndex]
	sortedNumbers = append(sortedNumbers[:middleIndex], sortedNumbers[middleIndex+1:]...)
	minors, majors := partitioner(sortedNumbers, index)
	return append(append(
		quicksort(minors), index),
		quicksort(majors)...)
}

func partitioner(numbers []int, index int) (minors []int, majors []int) {
	for _, n := range numbers {
		if n <= index {
			minors = append(minors, n)
		} else {
			majors = append(majors, n)
		}
	}
	return minors, majors
}
