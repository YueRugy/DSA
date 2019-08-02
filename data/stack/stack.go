package stack

import (
	"fmt"
)

type Stack struct {
	data  int
	pNext *Stack
}

func InitStack() *Stack {
	stack := new(Stack)
	return stack
}

func Push(stack *Stack, val int) {
	if stack != nil {
		pTemp := stack.pNext
		node := new(Stack)
		node.data = val
		node.pNext = pTemp
		stack.pNext = node
	}
}

func IsEmpty(stack *Stack) bool{
	return stack.pNext == nil
}

func Pop(stack *Stack) (bool, int) {
	if stack == nil || IsEmpty(stack) {
		return false, -1
	}
	temp := stack.pNext
	stack.pNext = stack.pNext.pNext
	return true, temp.data
}
func Iter(stack *Stack) {
	if stack != nil && !IsEmpty(stack) {
		for temp := stack.pNext; temp != nil; temp = temp.pNext {
			fmt.Printf("val==%d  ", temp.data)
		}
		fmt.Println()
	}
}
