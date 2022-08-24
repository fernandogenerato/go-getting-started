package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go run(c)

	for value := range c {
		fmt.Println(value)
	}
}

func run(c chan<- int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
