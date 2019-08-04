package main

import (
	_ "dsa/linkedList/arrayStack"
	_ "dsa/linkedList/cycle"
	_ "dsa/linkedList/doubleLinked"
	_ "dsa/linkedList/linked"
	"dsa/linkedList/linkedStack"
	_ "fmt"
)

func main() {
	ls := linkedStack.InitLinkedStack()
	for index := 1; index <= 6; index++ {
		linkedStack.Push(ls, index)
	}
	linkedStack.Iter(ls)
	for index := 7; index >= 0; index-- {
		linkedStack.Pop(ls)
	}
	linkedStack.Iter(ls)
	//as := arrayStack.InitArrayStack()
	//for index := 1; index <= 6; index++ {
	//	arrayStack.Push(as, index)
	//}
	//arrayStack.Iter(as)
	//for index := 6; index >= 0; index-- {
	//	arrayStack.Pop(as)
	//}
	//arrayStack.Iter(as)

	//pCycle := cycle.InitCycle(5)
	//cycle.Iter(pCycle)
	//cycle.CountNumber(5, 2, 3)

	//pHead := doubleLinked.InitDoubleLinked()
	//doubleLinked.Append(pHead, 1)
	//doubleLinked.Append(pHead, 2)
	//doubleLinked.Append(pHead, 3)
	//doubleLinked.AddPre(pHead, 6)
	//doubleLinked.Add(pHead, 7, 1)
	//doubleLinked.Add(pHead, 8, 6)
	//doubleLinked.Add(pHead, 5, 3)
	//doubleLinked.Delete(pHead, 2)
	//doubleLinked.Iter(pHead)
	//doubleLinked.AddPre(pHead, 1)

	//doubleLinked.AddPre(pHead, 2)
	//doubleLinked.AddPre(pHead, 3)
	//doubleLinked.Iter(pHead)

	//ll := linked.InitList()
	//linked.AddFive(ll)
	//linked.Iter(ll)
	//linked.Sort(ll)
	//linked.Iter(ll)
	//linked.Reverse(&ll)
	//linked.Iter(ll)
	//linked.ReverseShow(ll)
	//linked.Iter(ll)
	//linked.Reverse(&ll)
	//linked.Iter(ll)

}
