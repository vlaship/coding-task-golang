package problem0120

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTotalV2(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{"example1", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{"example2", [][]int{{-10}}, -10},
		{"example3", [][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}}, 14},
		{"example4", [][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}, {5, 6, 7, 8, 9}}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumTotalV2(tt.triangle)
			require.Equal(t, tt.expected, result)
		})
	}
}
