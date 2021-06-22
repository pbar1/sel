package sel

import (
	"reflect"
	"testing"
)

func TestFloydRivest(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Ascending",
			args{
				list: []int{0, 1, 2},
				k:    0,
			},
			[]int{0},
		},
		{
			"Descending",
			args{
				list: []int{2, 1, 0},
				k:    0,
			},
			[]int{0},
		},
		{
			"Random",
			args{
				list: []int{21, 2, 11, 14, 5},
				k:    1,
			},
			[]int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloydRivest(tt.args.list, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FloydRivest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFloydRivest(b *testing.B) {
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
				FloydRivest(input, bm.k)
			}
		})
	}
}
