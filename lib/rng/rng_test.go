package rng

import (
	"log"
	"testing"
)

func TestRandomDifferentInts(t *testing.T) {
	// manual dirty test
	log.Println(RandomDifferentInts(4, 0, 10))
}
