package crossover

import (
	"math/rand"
	"time"
)

func UniformCrossover(p1, p2 []int) []int {
	length := len(p1)
	child := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		coinFlip := rand.Intn(2)
		if coinFlip == 0 {
			child[i] = p1[i]
		} else {
			child[i] = p2[i]
		}
	}
	return child
}
