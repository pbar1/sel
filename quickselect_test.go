package sel

import (
	"fmt"
	"testing"
)

func TestQuickselect(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickselectInts(tt.args.list, tt.args.k); got != tt.want {
				t.Errorf("Quickselect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickselect(b *testing.B) {
	for _, bm := range benchmarks {
		for _, n := range ns {
			input := bm.genFunc(n)
			name := fmt.Sprintf("%s/%d", bm.name, n)
			b.Run(name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					QuickselectInts(input, bm.k)
				}
			})
		}
	}
}
