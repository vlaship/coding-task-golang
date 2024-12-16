package problem_1004

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestOnesV2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"example1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"example2", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestOnesV2(tt.nums, tt.k)
			require.Equal(t, tt.expected, result)
		})
	}
}
