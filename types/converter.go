package main

import "fmt"

type List []string

func main() {

	list := List{"Meat", "Fish", "Chicken"}
	slice := []string{"Meat", "Fish", "Chicken"}

	printList(list)
	printSlice(slice)

}

func printSlice(slice []string) {
	fmt.Println("Slice :", slice)
}

func printList(list List) {
	fmt.Println("Shopping List :", list)
}
