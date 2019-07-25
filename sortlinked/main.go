package main

import (
	"fmt"
)

type heroLinked struct {
	head *hero
	size int
	tail *hero
}

func InitHeroLinked() *heroLinked {
	h1 := heroLinked{
		head: nil,
		tail: nil,
		size: 0,
	}
	return &h1
}

type hero struct {
	name     string
	no       int
	nickname string
	next     *hero
}

func (h *heroLinked) addTail(h1 *hero) {
	//尾结点的next指向新加入的
	h.tail.next = h1
	//重新定义尾结点
	h.tail = h1
	//length++
	h.size++
}
func (h *heroLinked) add(h1 *hero) {
	if h.size == 0 { //链表有0个元素直接在链表末尾添加
		h.head = h1
		h.tail = h1
		h.size++
	} else if h.size == 1 { //有一个元素是比较no决定谁是tail 即可
		if h.tail.no < h1.no {
			h.addTail(h1)
		} else {
			h.head = h1
			h1.next = h.tail
			h.size++
		}
	} else {
		if h.head.no > h1.no {
			h.addHead(h1)
		} else if h.tail.no < h1.no {
			h.addTail(h1)
		} else {
			for front := h.head; front != nil; {
				if front.next != nil && front.next.no > h1.no {

					h1.next = front.next
					front.next = h1
					h.size++
					break
				}
				front = front.next

			}
		}
	}
}
func (h *heroLinked) addHead(h1 *hero) {

	h1.next = h.head
	h.head = h1
	h.size++
}
func (h *heroLinked) removeHead() {
	if h.head == h.tail {
		h.head = nil
		h.tail = nil
	} else {
		h.head = h.head.next
	}
	h.size--
}

func (h *heroLinked) removeTail() {
	if h.head == h.tail {
		h.removeHead()
	} else {
		for front := h.head.next; ; {
			if front.next.next == nil {
				front.next = nil
				h.tail = front
				h.size--
			}
		}
	}
}

func (h *heroLinked) show() {
	if h.size > 0 {
		for front := h.head; front != nil; {
			fmt.Println(front)
			front = front.next
		}
	}
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
	linked.removeHead()
	linked.add(&hero3)
	linked.add(&hero1)
	linked.add(&hero2)
	linked.show()

}
