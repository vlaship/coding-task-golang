package problem_3264

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFinalArray(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		k          int
		multiplier int
		want       []int
	}{
		{
			name:       "Example 1: Multiple operations with different values",
			nums:       []int{2, 1, 3, 5, 6},
			k:          5,
			multiplier: 2,
			want:       []int{8, 4, 6, 5, 6},
		},
		{
			name:       "Example 2: Small array with multiple operations",
			nums:       []int{1, 2},
			k:          3,
			multiplier: 4,
			want:       []int{16, 8},
		},
		{
			name:       "Single element array",
			nums:       []int{5},
			k:          2,
			multiplier: 3,
			want:       []int{45},
		},
		{
			name:       "No operations (k=0)",
			nums:       []int{1, 2, 3},
			k:          0,
			multiplier: 2,
			want:       []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFinalState(tt.nums, tt.k, tt.multiplier)
			require.Equal(t, tt.want, got)
		})
	}
}
