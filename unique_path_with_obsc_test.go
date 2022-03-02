package algors

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			},
			want: 2,
		},
		{
			name: "test-1",
			args: args{
				obstacleGrid: [][]int{{0, 1}, {0, 0}},
			},
			want: 1,
		},
		{
			name: "test-2",
			args: args{
				obstacleGrid: [][]int{{1, 0}},
			},
			want: 0,
		},
		{
			name: "test-3",
			args: args{
				obstacleGrid: [][]int{{0, 0}, {1, 1}, {0, 0}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
