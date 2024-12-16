package problem_2762

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContinuousSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{
			name: "Example 1",
			nums: []int{5, 4, 2, 4},
			want: 8,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "All same numbers",
			nums: []int{2, 2, 2, 2},
			want: 10,
		},
		{
			name: "Exactly difference of 2",
			nums: []int{1, 3, 1},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := continuousSubarrays(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
