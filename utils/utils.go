package utils

import (
	"math/rand"
	"sync"
	"time"
)

var newRand rand.Rand
var once sync.Once

func GetRandom() rand.Rand {
	once.Do(func() {
		seed := time.Now().Unix()
		source := rand.NewSource(seed)
		newRand = rand.New(source)
	})

	return newRand
}
