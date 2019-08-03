package main

import (
	"dsa/linkedList/doubleLinked"
	_ "dsa/linkedList/linked"
	_ "fmt"
)

func main() {

	pHead := doubleLinked.InitDoubleLinked()
	doubleLinked.Append(pHead, 1)
	doubleLinked.Append(pHead, 2)
	doubleLinked.Append(pHead, 3)
	doubleLinked.AddPre(pHead, 6)
	doubleLinked.Add(pHead, 7, 1)
	doubleLinked.Add(pHead, 8, 6)
	doubleLinked.Add(pHead, 5, 3)
	doubleLinked.Delete(pHead, 2)
	doubleLinked.Iter(pHead)
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
