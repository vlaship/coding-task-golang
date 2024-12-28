package problem_0689

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestMaxSumOfThreeSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 2, 1, 2, 6, 7, 5, 1},
			k:        2,
			expected: []int{0, 3, 5},
		},
		{
			name:     "example 2",
			nums:     []int{1, 2, 1, 2, 1, 2, 1, 2, 1},
			k:        2,
			expected: []int{0, 2, 4},
		},
		{
			name:     "single element subarrays",
			nums:     []int{4, 3, 2, 1, 5, 6, 7, 8},
			k:        1,
			expected: []int{5, 6, 7},
		},
		{
			name:     "equal sums but lexicographically smaller",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
			k:        2,
			expected: []int{0, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSumOfThreeSubarrays(tt.nums, tt.k)
			require.Equal(t, tt.expected, result)
		})
	}
}
