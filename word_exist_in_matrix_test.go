package algors

import "testing"

/*
* a b c
* e f g
* h i j
 */

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test-1",
		// 	args: args{
		// 		board: [][]byte{{'a', 'b', 'c'}, {'e', 'f', 'g'}, {'h', 'i', 'j'}},
		// 		word:  "bfij",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test-1",
		// 	args: args{
		// 		board: [][]byte{{'a', 'b', 'c'}, {'e', 'f', 'g'}, {'h', 'i', 'j'}},
		// 		word:  "afij",
		// 	},
		// 	want: false,
		// },
		{
			name: "test-3",
			args: args{
				board: [][]byte{{'a'}},
				word:  "ab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
