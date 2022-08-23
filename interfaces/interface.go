package main

import "fmt"

func main() {
	list := GenericList{1, "GoLang", 100, true, 3.14, "A", false}
	fmt.Println("Origin List :", list)
	list.removeEnd()
	fmt.Println("List without end (item) :", list)
	list.removeStart()
	fmt.Println("List without start (item) :", list)
	list.removeIndex(len(list) / 2)
	fmt.Println("List without middle (item) :", list)
}

type GenericList []interface{}

func (list *GenericList) removeEnd() interface{} {
	return list.removeIndex(len(*list) - 1)
}
func (list *GenericList) removeStart() interface{} {
	return list.removeIndex(0)
}
func (list *GenericList) removeIndex(index int) interface{} {
	l := *list
	removed := l[index]
	*list = append(l[0:index], l[index+1:]...)
	return removed
}

/*
	Result example :
		Origin List : [1 GoLang 100 true 3.14 A false]
		List without end (item) : [1 GoLang 100 true 3.14 A]
		List without start (item) : [GoLang 100 true 3.14 A]
		List without middle (item) : [GoLang 100 3.14 A]
*/
