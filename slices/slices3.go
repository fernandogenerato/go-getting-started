package main

import "fmt"

//cut off elements
func main() {

	//remove 00
	arr := []int{00, 10, 20, 30, 40, 50, 60}
	arr = arr[1:]
	fmt.Println(arr)

	//remove 30 & 40
	arr = append(arr[:2], arr[4:]...)
	fmt.Println(arr)

}
