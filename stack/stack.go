package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Stack created with size", stack.Size())
	fmt.Println("Stack is empty?", stack.Empty())

	stack.Push("Go")
	stack.Push(2009)
	stack.Push(3.14)
	stack.Push("End")
	fmt.Println("Size after push 4 values",
		stack.Size())
	fmt.Println("Is empty?", stack.Empty())

	for !stack.Empty() {
		v, _ := stack.Pop()
		fmt.Println("Unstacked", v)
		fmt.Println("Size: ", stack.Size())
		fmt.Println("Is empty?", stack.Empty())
	}
	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}
func (stack Stack) Empty() bool {
	return stack.Size() == 0
}
func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Stack is empty!")
	}
	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}
