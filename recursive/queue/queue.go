package queue

import (
	"fmt"
	"math"
)

var array = make([]int, 8)
var max = len(array)
var methodCount = 0
var count = 0

func Queue() {
	goQueue(0)
	fmt.Println(count)
	fmt.Println(methodCount)

}
func goQueue(n int) {
	if n >= max {
		count++
		iter()
		return
	}
	for index := 0; index < max; index++ {
		array[n] = index
		if judge(n) {
			goQueue(n + 1)
		}
	}
}

func judge(n int) bool {
	methodCount++
	if n > max {
		return false
	}
	for index := 0; index < n; index++ {
		if array[index] == array[n] || math.Abs(float64(array[index]-array[n])) == math.Abs(float64(index-n)) {
			return false
		}
	}
	return true
}

func iter() {
	for index := 0; index < max; index++ {
		fmt.Printf("%d  ", array[index])
	}
}
