package main

import (
	"fmt"
	"regexp"
	"strings"
)

//anonymous function example
func anonFunc() {
	expr := regexp.MustCompile("\\b\\w")
	text := "antonio carlos jobim"
	transform := func(t string) string { return strings.ToUpper(t) }
	processed := expr.ReplaceAllStringFunc(text, transform)

	fmt.Println(transform(text))
	fmt.Println(processed)
}

func simpleRegexReplace() {
	text := "Golang has 13 years"
	expr := regexp.MustCompile("\\d")
	fmt.Println(expr.ReplaceAllString(text, "YY"))
}
