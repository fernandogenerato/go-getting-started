package main

import "fmt"

func main() {
	values := []int{3, -2, 5, 7, 8, 22, 32, -1}
	fmt.Println(CalculateSum(values))
	fmt.Println(CalculateProduct(values))
}

type Aggregator func(n, m int) int

func Aggregate(values []int, initial int, a Aggregator) int {
	value := initial
	for _, v := range values {
		value = a(v, value)
	}
	return value
}

func CalculateSum(values []int) int {
	sum := func(n, m int) int {
		return n + m
	}
	return Aggregate(values, 0, sum)
}

func CalculateProduct(values []int) int {

	multiplication := func(n, m int) int {
		return n * m
	}
	return Aggregate(values, 1, multiplication)
}
