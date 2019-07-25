package main

import (
	"errors"
	"fmt"
	"os"
)

type singleQueue struct {
	array   [5]int
	front   int
	re      int
	maxSize int
}

func CreateDefalultSingleQueue() singleQueue {
	sq := singleQueue{
		front:   -1,
		re:      -1,
		maxSize: 5,
	}
	return sq
}
func (this *singleQueue) add(num int) error {
	if this.re >= this.maxSize-1 {
		fmt.Println("queue full")
		return errors.New("queue full")
	}
	this.re++
	this.array[this.re] = num
	return nil
}
func (this *singleQueue) show() {
	for i := this.front + 1; i <= this.re; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}
func (this *singleQueue) get() (int, error) {
	if this.front == this.re {
		fmt.Println("queue empty")
		return -1, errors.New("queue empty")
	}
	this.front++
	return this.array[this.front], nil
}

func main() {
	value := CreateDefalultSingleQueue()
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
