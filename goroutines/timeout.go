package main

import (
	"fmt"
	"time"
)

//waiting gouroutine by 5 sec and validate timeout marked as 2 sec
func main() {
	cn := make(chan bool, 1)
	go ex(cn)
	fmt.Println("waiting...")
	validateTimeout(cn)
}

func validateTimeout(cn chan bool) {
	end := false
	for !end {
		select {
		case end = <-cn:
			fmt.Println("End!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
		}

	}
}

func ex(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}
