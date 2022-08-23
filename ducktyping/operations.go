package main

import (
	"fmt"
	"time"
)

type Operation interface {
	Calculate() int
}

type Age struct {
	birthDate int
}

func (a Age) Calculate() int {
	return time.Now().Year() - a.birthDate
}
func (a Age) ToString() string {
	return fmt.Sprintf("Age since of %d", a.birthDate)
}

func acumulate(operations []Operation) int {
	acumulator := 0

	for _, op := range operations {
		value := op.Calculate()
		fmt.Printf("%v = %d\n", op, value)
		acumulator += value
	}
	return acumulator
}

func main() {

	ages := make([]Operation, 3)
	ages[0] = Age{birthDate: 1993}
	ages[1] = Age{birthDate: 2002}
	ages[2] = Age{birthDate: 1985}

	sum := acumulate(ages)

	fmt.Printf("age's sum : %d", sum)

}
