package algors

import (
	"testing"
)

// [3,9,20,15,7]
// [9,3,15,20,7]

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildTree(tt.args.preorder, tt.args.inorder)
		})
	}
}
