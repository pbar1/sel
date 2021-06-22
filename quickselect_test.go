package sel

import (
	"reflect"
	"testing"
)

func TestQuickselect(t *testing.T) {
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
			if got := Quickselect(tt.args.list, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quickselect() = %v, want %v", got, tt.want)
			}
		})
	}
}
