package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: conversor <values> <unit>")
		os.Exit(1)
	}

	argsLen := len(os.Args) - 1
	unitOrigin := os.Args[argsLen]
	originValues := os.Args[1:argsLen]
	var unitDestiny string
	if unitOrigin == "celsius" {
		unitOrigin = "fahrenheit"
	} else if unitOrigin == "quilometers" {
		unitDestiny = "miles"
	} else {
		fmt.Printf("%s it's not a known unit!", unitDestiny)
		os.Exit(1)
	}

	for index, value := range originValues {
		originValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("The value %s at position %d "+"is not a valid number!\n", value, index)
			os.Exit(1)
		}
		var valueDestiny float64
		if unitOrigin == "celsius" {
			valueDestiny = originValue*1.8 + 32
		} else {
			valueDestiny = originValue / 1.60934
		}
		fmt.Printf("%.2f %s = %.2f %s\n",
			originValue, unitOrigin,
			valueDestiny, unitDestiny)
	}

}
