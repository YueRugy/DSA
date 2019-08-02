package main

import "fmt"

//定义结构体 数据域 指针域
type Node struct {
	data  int
	pNext *Node
}

//链表的遍历
func traverseLinked(pHead *Node) {
	if nil == pHead {
		fmt.Println("不存在这个链表")
		return
	}
	headData := (*pHead).pNext
	for headData != nil {
		fmt.Printf("val==%d\t", (*headData).data)
		headData = (*headData).pNext
	}
	fmt.Println()
}

//创建一个链表
func createLinked() *Node {
	//创建一个不存放数据存放指向实际意义的第一个节点第一个的头结点
	var headNode *Node = new(Node)
	pTail := headNode //便于链表的添加设置一个尾结点
	for i := 1; i < 6; i++ {
		pTemp := new(Node)
		(*pTemp).data = i
		(*pTemp).pNext = nil
		(*pTail).pNext = pTemp
		pTail = pTemp
	}
	return headNode
}
func lengthLinked(pHead *Node) int {
	if pHead == nil {
		fmt.Println("传入指针不能为空")
	}
	pData := pHead.pNext
	length := 0
	if pData == nil {
		return 0
	}
	for temp := pData; temp != nil; temp = temp.pNext {
		length++
	}
	return length
}

//pos 表示从第几个位置插入从1开始
func insert(pHead *Node, pos int, val int) bool {
	if pHead == nil {
		return false
	}
	index := 0
	temp := pHead
	for ; temp != nil && index < pos-1; index++ {
		temp = (*temp).pNext
	}
	if index != pos-1 {
		return false
	}
	node := new(Node)
	(*node).data = val
	pTemp := temp.pNext
	temp.pNext = node
	(*node).pNext = pTemp
	return true
}

func soft(pHead *Node) {
	if pHead == nil {
		fmt.Println("链表不存在")
	}
	length := lengthLinked(pHead)
	if length == 0 {
		return
	}
	for i, pData := 0, pHead.pNext; i < length-1; i, pData = i+1, pData.pNext {
		var address *Node = pData
		for j, temp := i+1, pData.pNext; j < length; j, temp = j+1, temp.pNext {
			if (*address).data < (*temp).data {
				address = temp
			}
		}
		if address != pData {
			tempData := (*pData).data
			(*pData).data = (*address).data
			(*address).data = tempData
		}
	}
}

//pos 从1开始
func delete(pHead *Node, pos int) bool {
	if pHead == nil {
		return false
	}

	length := lengthLinked(pHead)
	if pos > length {
		return false
	}
	index := 0
	temp := pHead
	for ; index < pos-1 && temp != nil; index++ {
		temp = (*temp).pNext
	}
	if index != pos-1 {
		return false
	}
	temp.pNext = temp.pNext.pNext
	return true
}

func reverseLinked(pHead **Node) {
	if *pHead == nil {
		return
	}

	newHead := new(Node)
	for pTemp := (*pHead).pNext; pTemp != nil; pTemp = pTemp.pNext {
		node := new(Node)
		node.data = pTemp.data

		if newHead.pNext == nil {
			newHead.pNext = node
		} else {
			pNode := newHead.pNext
			newHead.pNext = node
			node.pNext = pNode
		}
	}
	*pHead = newHead

}

func main() {
	pHead := createLinked()
	traverseLinked(pHead)
	//delete(pHead, 3)
	reverseLinked(&pHead)
	traverseLinked(pHead)

	//	fmt.Println(lengthLinked(pHead))
	//	soft(pHead)
	//	traverseLinked(pHead)
	//	insert(pHead, 1, 35)
	//	traverseLinked(pHead)
}
