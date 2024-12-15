package problem0121

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"example1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"example2", []int{7, 6, 4, 3, 1}, 0},
		{"example3", []int{1}, 0},
		{"example4", []int{1, 2}, 1},
		{"example5", []int{1, 4, 2}, 3},
		{"example6", []int{2, 1, 4}, 3},
		{"example7", []int{2, 1, 2, 0, 1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			require.Equal(t, tt.expected, result)
		})
	}
}
