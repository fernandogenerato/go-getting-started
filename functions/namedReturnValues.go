package main

import "fmt"

func main() {

	dolarValue, realValue := FinalPrice(525.00)
	fmt.Printf("dolar value %.2f\n"+"dolar value %.2f\n", dolarValue, realValue)
}
func FinalPrice(priceValue float64) (dolarPrice float64, realPrice float64) {
	factor := 1.33
	taxConversion := 2.34
	dolarPrice = priceValue * factor
	realPrice = dolarPrice * taxConversion
	return dolarPrice, realPrice
}
