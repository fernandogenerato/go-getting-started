package main

import "fmt"

//ShoppingList defines a specific types
type ShoppingList []string

func main() {

	list := ShoppingList{
		"Lettuce",
		"Oil",
		"Meat",
		"Chicken",
		"Fish",
		"Chocolate"}

	protein, vegetables, others := ShoppingList.sort(list)

	fmt.Println("Proteins : ", protein)
	fmt.Println("Vegetables : ", vegetables)
	fmt.Println("Others : ", others)
}

//function for specific usage with ShoppingList
func (list ShoppingList) sort() ([]string, []string, []string) {
	var proteins, vegetables, others []string
	for _, item := range list {
		switch item {
		case "Lettuce":
			vegetables = append(vegetables, item)
		case "Meat", "Fish", "Chicken":
			proteins = append(proteins, item)
		default:
			others = append(others, item)
		}
	}
	return proteins, vegetables, others
}
