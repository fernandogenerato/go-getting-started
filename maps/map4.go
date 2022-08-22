package main

import (
	"fmt"
	"sort"
)

//sorted map
func main() {

	squares := make(map[int]int, 15)

	for index := 1; index <= 15; index++ {
		squares[index] = index * index
	}
	numbers := make([]int, 0, len(squares))
	for num := range squares {
		numbers = append(numbers, num)
	}
	sort.Ints(numbers)

	for _, number := range numbers {
		square := squares[number]
		fmt.Printf("%dˆ2 = %d\n", number, square)
	}
}
