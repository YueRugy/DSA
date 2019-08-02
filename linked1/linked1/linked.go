package linked1

import (
	"../utils"
	"fmt"
)

type Linked struct {
	data  int
	pNext *Linked
}

func InitLinked() *Linked {
	pHead := new(Linked)
	return pHead
}

//创建length==5的链表
func AddFive(pHead *Linked) {
	pTemp := pHead
	newRand := utils.GetRandom()
	fmt.Println(newRand)
	for index := 0; i < 5; i++ {
		fmt.Println()
	}
}
