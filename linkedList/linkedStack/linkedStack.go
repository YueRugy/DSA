package linkedStack

import (
	"fmt"
)

type LinkedStack struct {
	data int
	next *LinkedStack
}

func InitLinkedStack() *LinkedStack {
	ls := new(LinkedStack)
	return ls
}
func IsEmpty(ls *LinkedStack) bool {
	return ls.next == nil
}
func Push(ls *LinkedStack, val int) {
	if ls != nil {
		node := new(LinkedStack)
		node.data = val
		node.next = ls.next
		ls.next = node
	}
}
func Pop(ls *LinkedStack) (bool, int) {
	if ls == nil || IsEmpty(ls) {
		return false, -1
	}
	val := ls.next.data
	ls.next = ls.next.next
	return true, val
}
func Iter(ls *LinkedStack) {
	if ls != nil && !IsEmpty(ls) {
		for temp := ls.next; temp != nil; temp = temp.next {
			fmt.Printf("%d  ", temp.data)
		}
		fmt.Println()
	}
}
