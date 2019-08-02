package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func GetInt(num int) int {
	once.Do(func() {
		rand.Seed(time.Now().Unix())
		fmt.Println("execute init Seed")
	})
	number := rand.Intn(num) + 1
	return number
}
