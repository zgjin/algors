package algors

import "testing"

func Test_cuttingRope(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				n: 5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
