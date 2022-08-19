package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Loren", "ipsum", "lolor", "sit", "amet", "iua", "velit"}
	statistics := getStatistics(words)
	print(statistics)
}

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		count, found := statistics[initial]
		if found {
			statistics[initial] = count + 1
		} else {
			statistics[initial] = 1
		}
	}
	return statistics
}

func print(statistics map[string]int) {
	fmt.Println("count each initial word from string array :")
	for init, count := range statistics {
		fmt.Printf("%s = %d\n", init, count)
	}
}
