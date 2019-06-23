package boolGenerator

import (
	"math/rand"
	"time"
)

func New(time time.Time) bool {
	rand.Seed(time.UnixNano())
	var number = rand.Int()
	return number%2 == 0
}

func NewByPercent(time time.Time, hitPercent int) bool{
	rand.Seed(time.UnixNano())
	var number = rand.Intn(100)
	return number <= hitPercent;
}