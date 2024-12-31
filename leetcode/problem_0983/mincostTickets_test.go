package problem_0983

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinCostTickets(t *testing.T) {
	tests := []struct {
		name     string
		days     []int
		costs    []int
		expected int
	}{
		{
			name:     "Example 1",
			days:     []int{1, 4, 6, 7, 8, 20},
			costs:    []int{2, 7, 15},
			expected: 11,
		},
		{
			name:     "Example 2",
			days:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			costs:    []int{2, 7, 15},
			expected: 17,
		},
		{
			name:     "Single Day",
			days:     []int{1},
			costs:    []int{2, 7, 15},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mincostTickets(tt.days, tt.costs)
			require.Equal(t, tt.expected, got)
		})
	}
}
