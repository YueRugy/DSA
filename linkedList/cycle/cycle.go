package cycle

import (
	"dsa/linkedList/linked"
	"fmt"
)

type Cycle struct {
	data int
	next *Cycle
}

func InitCycle(num int) *Cycle {
	cycle := new(Cycle)
	cur := cycle
	for index := 1; index <= num; index++ {
		if index == 1 {
			cur.next = cycle
			cur.data = index
		} else {
			node := new(Cycle)
			cur.next = node
			node.data = index
			node.next = cycle
			cur = node
		}
	}
	return cycle
}
func Iter(cycle *Cycle) {
	if cycle != nil {
		for temp := cycle; true; temp = temp.next {
			fmt.Printf(" %d  ", temp.data)
			if temp.next == cycle {
				break
			}
		}
		fmt.Println()
	}
}

//总共有多少数，从第几个开始数，每次数几个 模拟约瑟夫问题
func CountNumber(totalNum int, startNo int, count int) {
	cycle := InitCycle(totalNum)
	pre, cur := cycle, cycle
	//确定pre,cur
	for index := 1; index <= startNo; index++ {
		if index > 1 {
			if index == 2 {
				cur = cur.next
			} else {
				cur = cur.next
				pre = pre.next
			}
		}
	}
	if pre.next != cur {
		for pre.next != cur {
			pre = pre.next
		}
	}
	list := linked.InitList()
	for true {
		if pre == cur {
			linked.Append(list, cur.data)
			break
		}
		for index := 1; index <= count-1; index++ {
			pre = pre.next
			cur = cur.next
		}
		linked.Append(list, cur.data)
		cur = cur.next
		pre.next = cur
	}
	linked.Iter(list)

}
