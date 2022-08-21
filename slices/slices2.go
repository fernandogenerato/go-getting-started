package main

import "fmt"

//slices + append usage
func main() {

	values := make([]int, 1)
	values = append(values, 1, 4, 5)
	missingValues := []int{2, 3}
	fmt.Println(values)
	fmt.Println(missingValues)
	values = append(values[:2], append(missingValues, values[2:]...)...)
	fmt.Println(values)
}
