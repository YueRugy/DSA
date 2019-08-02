package linked

import (
	"dsa/data/stack"
	"dsa/utils"
	"fmt"
)

type LinkedList struct {
	data  int
	pNext *LinkedList
}

func InitList() *LinkedList {
	linked := new(LinkedList)
	return linked
}

func AddFive(pHead *LinkedList) {
	temp := pHead
	for temp.pNext != nil {
		temp = temp.pNext
	}

	for i := 0; i < 5; i++ {
		node := new(LinkedList)
		val := utils.GetInt(10)
		node.data = val
		temp.pNext = node
		temp = node
	}
}
func Iter(pHead *LinkedList) {
	if pHead == nil || pHead.pNext == nil {
		return
	}
	for temp := pHead.pNext; temp != nil; temp = temp.pNext {
		fmt.Printf("val == %d", temp.data)
	}
	fmt.Println()
}

func Length(pHead *LinkedList) int {
	if pHead == nil {
		return -1
	}
	if pHead.pNext == nil {
		return 0
	}
	length := 0
	for temp := pHead.pNext; temp != nil; temp = temp.pNext {
		length++
	}
	return length
}

func Sort(pHead *LinkedList) {
	if pHead == nil || pHead.pNext == nil {
		return
	}
	length := Length(pHead)
	for index, temp := 0, pHead.pNext; index < length-1; index, temp = index+1, temp.pNext {
		address := temp
		for j, tNext := index+1, temp.pNext; j < length; j, tNext = j+1, tNext.pNext {
			if address.data < tNext.data {
				address = tNext
			}
		}
		if address != temp {
			tempData := temp.data
			temp.data = address.data
			address.data = tempData
		}
	}

}

func Reverse(pHead **LinkedList) {
	if pHead == nil || *pHead == nil || (*pHead).pNext == nil || (*pHead).pNext.pNext == nil {
		return
	}
	newLinked := new(LinkedList)
	for temp := (*pHead).pNext; temp != nil; temp = temp.pNext {
		tNext := newLinked.pNext
		node := new(LinkedList)
		node.data = temp.data
		newLinked.pNext = node
		node.pNext = tNext
	}
	*pHead = newLinked
}
func ReverseShow(pHead *LinkedList) {
	if pHead != nil && pHead.pNext != nil {
		if pHead.pNext.pNext == nil {
			fmt.Println(pHead.pNext.data)
		} else {
			pStack := stack.InitStack()
			for temp := pHead.pNext; temp != nil; temp = temp.pNext {
				stack.Push(pStack, temp.data)
			}
			stack.Iter(pStack)
		}
	}
}
