package algors

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		a  []int
		lo int
		hi int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "sort1",
			args: args{
				a:  []int{2, 4, 1, 0, 8, 8, 5},
				lo: 0,
				hi: 6,
			},
			want: []int{0, 1, 2, 4, 5, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aux = make([]int, tt.args.hi+1)
			MergeSort(tt.args.a, tt.args.lo, tt.args.hi)
			t.Log(tt.args.a)
		})
	}
}

func TestMergeBU(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "sort1",
			args: args{
				a: []int{5, 2, 1, 0, 6, 3},
			},
			want: []int{0, 1, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeBU(tt.args.a)
			t.Log(tt.args.a)
		})
	}
}
