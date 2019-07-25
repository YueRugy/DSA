package main

import (
	"errors"
	"fmt"
)

type heroLinked struct {
	head *hero
	size int
	tail *hero
}

func InitHeroLinked() *heroLinked {
	headHero := hero{}
	h1 := heroLinked{
		head: &headHero,
		tail: &headHero,
		size: 0,
	}
	return &h1
}
func (h *heroLinked) add(h1 *hero) {
	//尾结点的next指向新加入的
	h.tail.next = h1
	//重新定义尾结点
	h.tail = h1
	//length++
	h.size++
}
func (h *heroLinked) remove() error {
	if h.size == 0 {
		fmt.Println("linked length==0")
		return errors.New("linked empty")
	}
	//找到最后一个节点的前一个节点

	for front := h.head; ; {
		if front.next.next == nil {
			front.next = nil
			h.tail = front
			h.size--
			break
		}
		front = front.next
	}
	return nil
}

func (h *heroLinked) show() {
	if h.size > 0 {
		for front := h.head.next; front != nil; {
			fmt.Println(front)
			front = front.next
		}
	}
}

type hero struct {
	name     string
	no       int
	nickname string
	next     *hero
}

func main() {
	hero1 := hero{
		name:     "宋江1",
		nickname: "及时雨1",
		no:       1,
	}
	hero2 := hero{
		name:     "宋江2",
		nickname: "及时雨2",
		no:       2,
	}
	hero3 := hero{
		name:     "宋江3",
		nickname: "及时雨3",
		no:       3,
	}

	linked := InitHeroLinked()
	linked.add(&hero1)
	linked.show()
	linked.remove()
	linked.add(&hero1)
	linked.add(&hero2)
	linked.add(&hero3)
	linked.show()
}
