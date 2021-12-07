package algors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		a  []int
		lo int
		hi int
	}
	tests := []struct {
		name   string
		args   args
		expect []int
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				a:  []int{5, 3, 2, 1},
				lo: 0,
				hi: 3,
			},
			expect: []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.a, tt.args.lo, tt.args.hi)
			assert.Equal(t, tt.expect, tt.args.a)
		})
	}
}
