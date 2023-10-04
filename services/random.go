package services

import (
	"math/rand"
)

type RandomGenerator struct{}

func (RandomGenerator) Intn(n int) int {
	return rand.Intn(n)
}
