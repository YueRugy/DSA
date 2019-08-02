package doubleLinked

import (
	_ "fmt"
)

type DoubleLinked struct {
	data int
	pre  *DoubleLinked
	next *DoubleLinked
}

func InitDoubleLinked() *DoubleLinked {
	pHead := new(DoubleLinked)
	return pHead
}
func AddPre(dl *DoubleLinked, val int) {
	if dl != nil {
		node := new(DoubleLinked)
		node.pre = dl
		node.next = dl.next
		dl.next = node
	}
}

func Append(dl *DoubleLinked, val int) {
	if dl != nil {
		for temp := dl.next; temp != nil; temp = temp.next {
			node := new(DoubleLinked)
			node.pre = temp
			temp.next = node
		}
	}
}

func Add(dl *DoubleLinked, val int, index int) {
	if dl != nil {
		tempIndex := 0
		for temp := dl.next; temp != nil; temp = temp.next {
			tempIndex++
			if tempIndex == index {
				break
			}
		}
		if tempIndex == index {
			node := new(DoubleLinked)
			node.pre = temp
			node.next = temp.next
			tempNode := temp.next
			temp.next = node
			if tempNode != nil {
				tempNode.pre = node
			}
		}
	}
}

func Delete(dl *DoubleLinked, index int) {
	if dl != nil {
		tempIndex := 0
		for temp := dl.next; temp != nil; temp = temp.next {
			tempIndex++
			if tempIndex == index {
				break
			}
		}
		if tempIndex == index {
			temp.pre.next = temp.next
			if temp.next != nil {
				temp.next.pre = temp.pre
			}
		}
	}
}
