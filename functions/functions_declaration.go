package main

import "fmt"

func main() {
	fmt.Println(soma(100, 200))
	imprimirDados("fernando", 29)
}
func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos.", nome, idade)
}
func soma(n1, n2 int) int {
	return n1 + n2
}
