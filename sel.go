package sel

import (
	"math/rand"
	"time"
)

var RNG *rand.Rand

func init() {
	RNG = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Greater(i, j int) bool
	Swap(i, j int)
	Get(i int) interface{}
}
