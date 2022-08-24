package main

import (
	"fmt"
	"time"
)

//example : goroutine only runs while main method is alive
func main() {

	go sleap()
	fmt.Println("Main sleeping 3 sec...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main terminated...")

}
func sleap() {
	fmt.Println("Goroutine sleeping 5 sec...")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine terminated...")

}

func print(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", n)
		time.Sleep(2000 * time.Millisecond)
	}
}
