package stack

import (
	"fmt"
)

type Node struct {
	data  int
	pNext *Node
}
type Stack struct {
	pTop    *Node
	pBottom *Node
}

func initStack() *Stack {
	stack := new(Stack)
	pHead := new(Node)
	pHead.pNext = nil
	stack.pBottom, stack.pTop = pHead, pHead
	return stack
}
func traverse(pStack *Stack) {
	if pStack != nil {
		for pTemp := pStack.pTop; pTemp != pStack.pBottom; pTemp = pTemp.pNext {
			fmt.Printf("%d  ", pTemp.data)
		}
		fmt.Println()
	}
}
func push(pStack *Stack, val int) {
	node := new(Node)
	node.data = val
	node.pNext = pStack.pTop
	pStack.pTop = node
}

func pop(pStack *Stack) (bool, int) {
	if pStack == nil || pStack.pBottom == pStack.pTop {
		return false, -1
	}
	val := pStack.pTop.data
	//pTemp := pStack.pTop
	pStack.pTop = pStack.pTop.pNext
	//val := pTemp.data
	//pTemp = nil //删除空间go 自动回收
	return true, val
}

func clear(pStack *Stack) {
	if pStack != nil {
		pStack.pTop = pStack.pBottom
	}
}

func main() {
	stack := initStack()
	push(stack, 1)
	push(stack, 2)
	push(stack, 3)
	push(stack, 4)
	push(stack, 5)
	//traverse(stack)
	//pop(stack)
	//traverse(stack)
	//pop(stack)
	//traverse(stack)
	//pop(stack)
	//traverse(stack)
	//pop(stack)
	//traverse(stack)
	//pop(stack)
	//traverse(stack)
	//pop(stack)
	//traverse(stack)
	clear(stack)
	traverse(stack)
}
