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

const (
	BoolType = iota
	IntType
	Int32Type
	Int8Type
	Int64Type
	RuneType
	StringType
	Float32Type
	Float64Type
)

func AssertType(t interface{}) int {

	switch t.(type) {
	case bool:
		return BoolType
	case rune:
		return RuneType
	case string:
		return StringType
	default:
		return -1
	}
}
