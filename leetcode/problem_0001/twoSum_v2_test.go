package problem_0001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSumV2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected [][]int
	}{
		{"example1", []int{1, 3, 3, 4, 5, 2}, 6, [][]int{{3, 3}, {1, 5}, {4, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSumV2(tt.input, tt.target)
			require.Equal(t, tt.expected, result)
		})
	}
}
