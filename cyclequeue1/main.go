package main

import (
	"errors"
	"fmt"
	"os"
)

type cycle struct {
	maxSize int
	head    int
	tail    int
	array   [6]int
}

func CreateDefaultCycle() cycle {
	return cycle{
		maxSize: 6,
		head:    0,
		tail:    0,
	}
}

func (queue *cycle) add(num int) error {
	if queue.isFull() {
		fmt.Println("queue full")
		return errors.New("queue full")
	}
	queue.array[queue.tail] = num
	queue.tail = (queue.tail + 1) % queue.maxSize
	return nil
}
func (queue *cycle) get() (int, error) {
	if queue.isEmpty() {
		fmt.Println("queue empty")
		return -1, errors.New("queue empty")
	}
	num := queue.array[queue.head]
	queue.head = (queue.head + 1) % queue.maxSize
	return num, nil
}
func (queue *cycle) getSize() int {
	return (queue.tail - queue.head + queue.maxSize) % queue.maxSize
}
func (queue *cycle) show() {
	for i, j := 0, queue.head; i < queue.getSize(); i++ {
		fmt.Printf("array[%d]==%d", j, queue.array[j])
		fmt.Println()
		j = (j + 1) % queue.maxSize
	}
}

func (queue *cycle) isFull() bool {
	return (queue.tail+1)%queue.maxSize == queue.head
}

func (queue *cycle) isEmpty() bool {
	return queue.tail == queue.head
}

func main() {
	value := CreateDefaultCycle()
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
