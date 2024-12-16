package problem_1760

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinimumSize(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		maxOperations int
		expected      int
	}{
		{
			name:          "Example 1",
			nums:          []int{9},
			maxOperations: 2,
			expected:      3,
		},
		{
			name:          "Example 2",
			nums:          []int{2, 4, 8, 2},
			maxOperations: 4,
			expected:      2,
		},
		{
			name:          "Single Bag No Operations",
			nums:          []int{5},
			maxOperations: 0,
			expected:      5,
		},
		{
			name:          "Multiple Bags No Operations",
			nums:          []int{1, 2, 3, 4, 5},
			maxOperations: 0,
			expected:      5,
		},
		{
			name:          "All Bags Same Size",
			nums:          []int{4, 4, 4, 4},
			maxOperations: 4,
			expected:      2,
		},
		{
			name:          "Large Numbers",
			nums:          []int{1000000000},
			maxOperations: 1000000000,
			expected:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumSize(tt.nums, tt.maxOperations)
			require.Equal(t, tt.expected, result)
		})
	}
}
