package rng

import (
	"math/rand"
)

func RandomDifferentInts(amount, min, max int) []int {
	nLen := max - min
	if nLen < amount {
		panic("RandomDifferentInts")
	}
	numbers := make([]int, nLen, nLen)
	for i := 0; i < nLen; i++ {
		numbers[i] = min + i
	}
	rand.Shuffle(nLen, func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	return numbers[0:amount]
}
