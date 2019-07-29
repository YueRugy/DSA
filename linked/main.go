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
		fmt.Printf("val==%d\n", (*headData).data)
		headData = (*headData).pNext
	}
}

//创建一个链表
func createLinked() *Node {
	//创建一个不存放数据存放指向实际意义的第一个节点第一个的头结点
	var headNode *Node = new(Node)
	pTail := headNode //便于链表的添加设置一个尾结点
	fmt.Println("请输入要创建的个数")
	var len int
	var val int
	fmt.Scanf("%d", &len)
	for i := 0; i < len; i++ {
		fmt.Println("请输入要增加的数字")
		fmt.Scanf("%d", &val)
		pTemp := new(Node)
		(*pTemp).data = val
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
func main() {
	pHead := createLinked()
	traverseLinked(pHead)
	fmt.Println(lengthLinked(pHead))
}
