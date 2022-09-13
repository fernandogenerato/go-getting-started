package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go write("hello world", c)

	for {
		// receive text and status of channel
		if text, isActive := <-c; isActive {
			fmt.Println(text)
		} else {
			break //condition to stop this for
		}
	}

	// resume of impl, iterate while chan is open. Avoid the infinite loop and verification of chan status.
	for t := range c {
		fmt.Println(t)
	}
}

func write(text string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- text // sending text to canal
		time.Sleep(time.Second)
	}

	close(c) // closing chan
}
