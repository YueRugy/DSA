package arrayStack

import (
	"fmt"
)

type ArrayStack struct {
	Top     int
	MaxSize int
	Array   [6]int
}

func InitArrayStack() *ArrayStack {
	as := new(ArrayStack)
	as.Top = -1
	array := [6]int{}
	as.MaxSize = len(as.Array)
	as.Array = array
	return as
}

func Length(as *ArrayStack) int {
	return as.Top + 1
}
func IsFull(as *ArrayStack) bool {
	return as.Top+1 == as.MaxSize
}

func IsEmpty(as *ArrayStack) bool {
	return as.Top == -1
}
func Push(as *ArrayStack, val int) {
	if as != nil {
		if !IsFull(as) {
			as.Top = as.Top + 1
			as.Array[as.Top] = val
		}
	}
}

func Pop(as *ArrayStack) (bool, int) {
	if as == nil || IsEmpty(as) {
		return false, -1
	}
	val := as.Array[as.Top]
	as.Top = as.Top - 1
	return true, val
}
func Iter(as *ArrayStack) {
	if as != nil && !IsEmpty(as) {
		for index := as.Top; index >= 0; index-- {
			fmt.Printf("%d  ", as.Array[index])
		}
		fmt.Println()
	}
}
