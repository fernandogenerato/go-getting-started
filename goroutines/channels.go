package main

import "fmt"

/*
create channel to stream int values
	c := make(chan int)
send value to channel
	c <- 33
receive value from channel
	value := <-c
*/
func main() {
	c := make(chan int)
	go run(c)
	value := <-c
	fmt.Println(value)

}

func r(c chan int) {
	c <- 100
}
