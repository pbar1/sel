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
}

type IntList []int

func (l IntList) Len() int              { return len(l) }
func (l IntList) Less(i, j int) bool    { return l[i] < l[j] }
func (l IntList) Greater(i, j int) bool { return l[i] > l[j] }
func (l IntList) Swap(i, j int)         { l[i], l[j] = l[j], l[i] }

type StringList []string

func (l StringList) Len() int              { return len(l) }
func (l StringList) Less(i, j int) bool    { return l[i] < l[j] }
func (l StringList) Greater(i, j int) bool { return l[i] > l[j] }
func (l StringList) Swap(i, j int)         { l[i], l[j] = l[j], l[i] }
