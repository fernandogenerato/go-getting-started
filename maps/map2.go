package main

import "fmt"

//map types of declaration
func main() {

	largeMap := map[string]string{
		"LX":  "Linux",
		"DOS": "Windows",
		"MAC": "MacOs"}

	fmt.Println(largeMap)
	fmt.Println(len(largeMap))

	sizedMap := make(map[string]string, 3000)
	fmt.Println(sizedMap)
	fmt.Println(len(sizedMap))

	capitals := map[string]string{
		"SP": "São Paulo",
		"PR": "Curitiba"}
	fmt.Println(capitals)
	capitals["RJ"] = "Rio de Janeiro"
	fmt.Println(capitals)

}
