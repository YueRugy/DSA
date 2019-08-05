package main

import (
	_ "dsa/linkedList/arrayStack"
	_ "dsa/linkedList/cal"
	_ "dsa/linkedList/cycle"
	_ "dsa/linkedList/doubleLinked"
	_ "dsa/linkedList/genericStack"
	_ "dsa/linkedList/linked"
	_ "dsa/linkedList/linkedStack"
	_ "dsa/linkedList/polandNotation"
	"fmt"
	_ "strings"
)

//func test(index ...*int) {
//	if index != nil {
//		*index[0] = 11
//	}
//}
func main() {

	var i int = 10
	num := i.(int)
	fmt.Println(num)

	//str := "3*5+20*(6+7*(8-6))"
	//num := cal.Cal(str)
	//fmt.Println(num)

	//str := "3,4,*,6,+,20,+"
	//num := polandNotation.CalStr(str)
	//fmt.Println(num)

	//var str string = "abcdefg"
	//array := []rune(str)
	//for _, temp := range array {
	//	fmt.Printf("%d  %T\n", temp, temp)
	//}
	//fmt.Println(string(array[0]))
	//gs := genericStack.InitGenericStack()
	//genericStack.Push(gs, "haha")
	//_, temp := genericStack.Pop(gs)
	//var num string = temp.(string)
	//fmt.Println(num)
	//index := 12
	//test(&index)
	//fmt.Println(index)
	//test()

	//ls := linkedStack.InitLinkedStack()
	//for index := 1; index <= 6; index++ {
	//	linkedStack.Push(ls, index)
	//}
	//linkedStack.Iter(ls)
	//for index := 7; index >= 0; index-- {
	//	linkedStack.Pop(ls)
	//}
	//linkedStack.Iter(ls)
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
