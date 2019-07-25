package main

import (
	"errors"
	"fmt"
	"os"
)

type cycleQueue struct {
	array   [5]int
	head    int
	tail    int
	maxSize int
}

func (this *cycleQueue) isFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}
func (this *cycleQueue) isEmpty() bool {
	return this.tail == this.head
}
func (this *cycleQueue) show() {
	fmt.Println(this.array)
	for i, j := 0, this.head; i < (this.tail-this.head+this.maxSize)%this.maxSize; i++ {
		//if j+1 == this.maxSize {
		//	i--
		//} else {
		fmt.Printf("array[%d]==%d", j, this.array[j])
		//}
		j = (j + 1) % this.maxSize
		fmt.Println()
	}
}

func (this *cycleQueue) add(num int) error {
	if this.isFull() {
		fmt.Println("queue full")
		return errors.New("queue full")
	}
	//if this.tail+1 == this.maxSize {
	//	this.tail = (this.tail + 1) % this.maxSize
	//	this.array[this.tail] = num
	//	this.tail = (this.tail + 1) % this.maxSize
	//} else {
	this.array[this.tail] = num
	this.tail = (this.tail + 1) % this.maxSize
	//}
	return nil
}
func (this *cycleQueue) get() (int, error) {
	if this.isEmpty() {
		fmt.Println("queue empty")
		return -1, errors.New("queue is empty")
	}
	num := 0
	//if this.head+1 == this.maxSize {
	//	this.head = (this.head + 1) % this.maxSize
	//	num = this.array[this.head]
	//	this.head = (this.head + 1) % this.maxSize
	//} else {
	num = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	//}
	return num, nil
}

func CreateDefalultCycleQueue() cycleQueue {
	sq := cycleQueue{
		head:    0,
		tail:    0,
		maxSize: 5,
	}
	return sq
}
func main() {
	value := CreateDefalultCycleQueue()
	queue := &value
	for {
		fmt.Println("1 add addqueue")
		fmt.Println("2 get getqueue")
		fmt.Println("3 show showqueue")
		fmt.Println("4 exit exit")
		var key string
		var val int
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入num")
			fmt.Scanln(&val)
			err := queue.add(val)
			if err != nil {
				err.Error()
			} else {
				fmt.Println("push success")
			}
		case "show":
			queue.show()
		case "get":
			num, err := queue.get()
			if err != nil {
				err.Error()
			} else {
				fmt.Printf("get from queue value is %d", num)
			}
		case "exit":
			os.Exit(0)
		}
	}
}
