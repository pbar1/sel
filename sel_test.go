package sel

import "math"

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

	ns = doubleRange(100, 100)

	benchmarks = []struct {
		name    string
		genFunc func(n int) []int
	}{
		// Random inputs
		{"Random", genRandom},
		{"Onezero", genOnezero},
		{"Twofaced", genTwofaced},
		// Deterministic inputs
		{"Sorted", genSorted},
		{"Rotated", genRotated},
		{"Organpipe", genOrganpipe},
		{"M3killer", genM3killer},
	}
)

// // BenchmarkSort is here as a control
// func BenchmarkSort(b *testing.B) {
// 	for _, bm := range benchmarks {
// 		for _, n := range ns {
// 			input := bm.genFunc(n)
// 			name := fmt.Sprintf("%s/%d", bm.name, n)
// 			b.Run(name, func(b *testing.B) {
// 				for i := 0; i < b.N; i++ {
// 					sort.Ints(input)
// 				}
// 			})
// 		}
// 	}
// }

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

// Musser’s “median-of-3 killer” sequence with n=4j and k=n/2
// Adapted from: https://programmingpraxis.com/2016/11/08/a-median-of-three-killer-sequence/#comment-61876
func genM3killer(n int) []int {
	list := make([]int, n)
	if n%2 != 0 {
		list[n-1] = n // fill in last element manually if n is odd
		n--           // make sure subsequent working value of n is always even
	}
	m := n / 2
	for i := 0; i < m; i++ {
		if i%2 == 0 {
			list[i] = i + 1 // first half, even indices
		} else {
			list[i] = m + i + m%2 // first half, odd indices
		}
		list[m+i] = (i + 1) * 2 // second half
	}
	return list
}

// Obtained by randomly permuting the elements of an m3killer sequence in
// positions noted by a:b and c:d.
func genTwofaced(n int) []int {
	list := genM3killer(n)
	a := 4 * int(math.Log2(float64(n)))
	b := n/2 - 1
	c := n/2 + a - 1
	d := n - 2
	slice1 := list[a : b+1]
	slice2 := list[c : d+1]
	RNG.Shuffle(len(slice1), func(i, j int) {
		slice1[i], slice1[j] = slice1[j], slice1[i]
	})
	RNG.Shuffle(len(slice2), func(i, j int) {
		slice2[i], slice2[j] = slice2[j], slice2[i]
	})
	return list
}
