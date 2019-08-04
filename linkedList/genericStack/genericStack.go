package genericStack

import (
	"fmt"
)

type GenericStack struct {
	data interface{}
	next *GenericStack
}

func Peek(gs *GenericStack) (bool, interface{}) {
	if gs == nil || gs.next == nil {
		return false, nil
	}

	temp := gs.next.data
	return true, temp
}

func InitGenericStack() *GenericStack {
	ls := new(GenericStack)
	return ls
}
func IsEmpty(ls *GenericStack) bool {
	return ls.next == nil
}
func Push(ls *GenericStack, val interface{}) {
	if ls != nil {
		node := new(GenericStack)
		node.data = val
		node.next = ls.next
		ls.next = node
	}
}
func Pop(ls *GenericStack) (bool, interface{}) {
	if ls == nil || IsEmpty(ls) {
		return false, nil
	}
	val := ls.next.data
	ls.next = ls.next.next
	return true, val
}
func Iter(ls *GenericStack) {
	if ls != nil && !IsEmpty(ls) {
		for temp := ls.next; temp != nil; temp = temp.next {
			fmt.Printf("%s  ", temp.data)
		}
		fmt.Println()
	}
}
