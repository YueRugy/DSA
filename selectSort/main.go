package main

import (
	"fmt"
)

func main() {
	var array = [6]int{23, 6, 89, 42, 7, 5}
	fmt.Println(array)
	soft(array[0:])
	fmt.Println(array)
}

//选择排序演示
func soft(array []int) {
	length := len(array)
	for i := 0; i < length-1; i++ { //执行length-1次
		var min int = i //最小值或最大值的默认下标是执行第几次的首个下表
		//设置min可以只交换一次
		for j := i + 1; j < length; j++ {
			if array[i] < array[j] {
				min = j
			}

			if array[i] < array[min] {
				temp := array[i]
				array[i] = array[min]
				array[min] = temp
			}
		}
	}
}
