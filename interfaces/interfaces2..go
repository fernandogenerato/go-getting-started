package main

import "fmt"

type Operation interface {
	Calculate() int
	Print(value int)
}

type Sum struct {
	op1, op2 int
}
type Sub struct {
	op1, op2 int
}

func (s Sum) Print(value int) {
	fmt.Printf("sum result: %d\n", value)
}
func (s Sub) Print(value int) {
	fmt.Printf("sub result: %d\n", value)
}

func (s Sum) Calculate() int {
	return s.op2 + s.op2
}

func (s Sub) Calculate() int {
	return s.op1 - s.op2
}

func (s Sum) String() string {
	return fmt.Sprintf("%d + %d", s.op1, s.op2)
}

func main() {

	var sum Operation
	sum = Sum{
		op1: 100,
		op2: 100,
	}

	sum.Print(sum.Calculate())

	var sub Operation
	sub = Sub{
		op1: 500,
		op2: 200,
	}

	sum.Print(sub.Calculate())

}
