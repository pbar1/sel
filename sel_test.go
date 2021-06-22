package sel

import (
	"fmt"
	"sort"
	"testing"
)

type args struct {
	list []int
	k    int
}

var (
	tests = []struct {
		name string
		args args
		want int
	}{
		{"Ascending", args{list: []int{0, 1, 2}, k: 0}, 0},
		{"Descending", args{list: []int{2, 1, 0}, k: 0}, 0},
		{"Random", args{list: []int{21, 2, 11, 14, 5}, k: 1}, 5},
	}

	ns = doubleRange(1_000_000, 16_000_000)

	benchmarks = []struct {
		name    string
		genFunc func(n int) []int
		k       int
	}{
		// Random inputs
		{"Random", genRandom, 1},
		{"Onezero", genOnezero, 1},
		// Deterministic inputs
		{"Sorted", genSorted, 1},
		{"Rotated", genRotated, 1},
		{"Organpipe", genOrganpipe, 1},
	}
)

// BenchmarkSort is here as a control
func BenchmarkSort(b *testing.B) {
	for _, bm := range benchmarks {
		for _, n := range ns {
			input := bm.genFunc(n)
			name := fmt.Sprintf("%s/%d", bm.name, n)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					sort.Ints(input)
				}
			})
		}
	}
}

func doubleRange(from, to int) []int {
	list := make([]int, 0)
	for i := from; i <= to; i += i {
		list = append(list, i)
	}
	return list
}

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

// The integers(1,2,...,n/2,n/2,...,2,1).
func genOrganpipe(n int) []int {
	list := make([]int, (n/2)*2)
	for i := 0; i < n/2; i++ {
		list[i] = i + 1
		list[(n/2)*2-1-i] = i + 1
	}
	return list
}

// TODO: m3killer

// TODO: twofaced

// choose calculates n choose k. Overflows are not detected, and Choose panics
// if n >= k >= 0 is violated.
func choose(n, k int) int {
	if k > n {
		panic("Choose: k > n")
	}
	if k < 0 {
		panic("Choose: k < 0")
	}
	if n <= 1 || k == 0 || n == k {
		return 1
	}
	if newK := n - k; newK < k {
		k = newK
	}
	if k == 1 {
		return n
	}
	// Our return value, and this allows us to skip the first iteration.
	ret := n - k + 1
	for i, j := ret+1, 2; j <= k; i, j = i+1, j+1 {
		ret = ret * i / j
	}
	return ret
}
