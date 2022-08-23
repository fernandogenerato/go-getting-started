package main

import "fmt"

type File struct {
	name  string
	size  float64
	chars int
	words int
	lines int
}

func main() {
	file := File{
		name:  "main.go",
		size:  3.14,
		chars: 120,
		words: 15,
		lines: 10,
	}
	fmt.Println(file)

	file2 := File{"main.java", 10, 12, 3, 5}
	fmt.Println(file2)

}
