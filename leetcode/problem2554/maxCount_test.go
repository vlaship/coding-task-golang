package problem2554

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxCount(t *testing.T) {
	tests := []struct {
		name     string
		banned   []int
		n        int
		maxSum   int
		expected int
	}{
		{
			name:     "Example 1",
			banned:   []int{1, 6, 5},
			n:        5,
			maxSum:   6,
			expected: 2,
		},
		{
			name:     "Example 2",
			banned:   []int{1, 2, 3, 4, 5, 6, 7},
			n:        8,
			maxSum:   1,
			expected: 0,
		},
		{
			name:     "Example 3",
			banned:   []int{11},
			n:        7,
			maxSum:   50,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxCount(tt.banned, tt.n, tt.maxSum)
			require.Equal(t, tt.expected, result)
		})
	}
}
