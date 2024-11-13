package problem0001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"example1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			require.Equal(t, tt.expected, result)
		})
	}
}
