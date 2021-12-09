package algors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_heap_sort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				a: []int{0, 1, 3, 8, 2, 4, 6, 9},
			},
			expected: []int{0, 1, 2, 3, 4, 6, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap_sort(tt.args.a)
			assert.Equal(t, tt.expected, tt.args.a)
		})
	}
}
