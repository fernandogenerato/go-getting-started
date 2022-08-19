package main

import "fmt"

func main() {

	var a [3]int
	numbers := [5]int{1, 2, 3, 4, 5}
	primes := [...]int{2, 3, 5, 7, 11, 13}
	names := [2]string{}

	fmt.Println(a, numbers, primes, names)

	fmt.Println(len(a))
	fmt.Println(len(numbers))
	fmt.Println(len(primes))
	println(len(names))

	var multA [2][2]int
	multA[0][0], multA[0][1] = 3, 5
	multA[1][0], multA[1][1] = 7, -2

	multB := [2][2]int{{2, 13}, {-1, 6}}

	fmt.Println("Multi A:", multA)

	fmt.Println("Multi A:", multB)

}
