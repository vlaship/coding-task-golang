package problem_0494

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTargetSumWays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
		{
			name:   "example 2",
			nums:   []int{1},
			target: 1,
			want:   1,
		},
		{
			name:   "zero target",
			nums:   []int{0, 0},
			target: 0,
			want:   4, // [+0+0], [+0-0], [-0+0], [-0-0]
		},
		{
			name:   "larger numbers",
			nums:   []int{1, 2, 3},
			target: 0,
			want:   2, // [+1-2+3], [-1+2-3]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTargetSumWays(tt.nums, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
