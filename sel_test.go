package sel

import (
	"sort"
	"testing"
)

// A random permutation of the integers 1 through n.
func genSorted(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}
	return list
}

// A random permutation of the integers 1 through n.
func genRandom(n int) []int {
	list := genSorted(n)
	RNG.Shuffle(n, func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

// A random permutation ceil(n/2) ones and floor(n/2) zeros.
func genOnezero(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = (i + 1) % 2
	}
	RNG.Shuffle(n, func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

// A sorted sequence rotated left once; i.e., (2,3,...,n,1).
func genRotated(n int) []int {
	list := make([]int, n)
	for i := 0; i < n-1; i++ {
		list[i] = i + 2
	}
	list[n-1] = 1
	return list
}

// TODO: Organpipe

// TODO: m3killer

// BenchmarkSort is here as a control
func BenchmarkSort(b *testing.B) {
	benchmarks := []struct {
		name    string
		genFunc func(n int) []int
		n       int
		k       int
	}{
		// Random inputs
		{
			"Random50K",
			genRandom,
			50000,
			1,
		},
		{
			"Random1M",
			genRandom,
			1000000,
			1,
		},
		{
			"Random16M",
			genRandom,
			16000000,
			1,
		},
		{
			"Onezero50K",
			genOnezero,
			50000,
			1,
		},
		{
			"Onezero1M",
			genOnezero,
			1000000,
			1,
		}, {
			"Onezero16M",
			genOnezero,
			16000000,
			1,
		},
		// Deterministic inputs
		{
			"Sorted50K",
			genSorted,
			50000,
			1,
		},
		{
			"Sorted1M",
			genSorted,
			1000000,
			1,
		},
		{
			"Sorted16M",
			genSorted,
			16000000,
			1,
		},
		{
			"Rotated50K",
			genRotated,
			50000,
			1,
		},
		{
			"Rotated1M",
			genRotated,
			1000000,
			1,
		}, {
			"Rotated16M",
			genRotated,
			16000000,
			1,
		},
	}
	for _, bm := range benchmarks {
		input := bm.genFunc(bm.n)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.Ints(input)
			}
		})
	}
}
