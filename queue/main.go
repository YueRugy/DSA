package main

import (
	"fmt"
	"os"
)

type Queue struct {
	front   int
	rear    int
	array   [6]int
	maxSize int
}

func initQueue() *Queue {
	queue := new(Queue)
	queue.rear, queue.front = 0, 0
	queue.array = [6]int{}
	queue.maxSize = len(queue.array)
	return queue
}

func empty(queue *Queue) bool {
	return queue.front == queue.rear
}

func full(queue *Queue) bool {
	return (queue.rear+1)%queue.maxSize == queue.front
}

func add(queue *Queue, val int) bool {
	if queue == nil || full(queue) {
		return false
	}
	queue.array[queue.rear] = val
	queue.rear = (queue.rear + 1) % queue.maxSize
	return true
}

func get(queue *Queue) (bool, int) {
	if queue == nil || empty(queue) {
		return false, -1
	}
	val := queue.array[queue.front]
	queue.front = (queue.front + 1) % queue.maxSize
	return true, val
}

func show(queue *Queue) {
	if queue != nil && !empty(queue) {

		for index := queue.front; index != queue.rear; {
			fmt.Printf("%d  ", queue.array[index])
			index = (index + 1) % queue.maxSize
		}
		fmt.Println()
	}
}

func length(queue *Queue) int {
	if queue == nil {
		return -1
	}
	return (queue.rear - queue.front + queue.maxSize) % queue.maxSize
}

func main() {
	queue := initQueue()
	for {
		fmt.Println("1 add addqueue")
		fmt.Println("2 get getqueue")
		fmt.Println("3 show showqueue")
		fmt.Println("4 exit exit")
		var key string
		var val int
		fmt.Scanln(&key)
		switch key {
		case "a":
			fmt.Println("请输入num")
			fmt.Scanln(&val)
			flag := add(queue, val)
			if flag {
				fmt.Println("push success")
			} else {
				fmt.Println("full or queue is nil")
			}
		case "s":
			show(queue)
		case "g":
			flag, num := get(queue)
			if !flag {
				fmt.Println("get failed queue is nil or queue is empty")
			} else {
				fmt.Printf("get from queue value is %d\n", num)
			}
		case "exit":
			os.Exit(0)
		}
	}
	//fmt.Println(*queue)
}
