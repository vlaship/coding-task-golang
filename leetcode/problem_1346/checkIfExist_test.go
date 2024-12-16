package problem_1346

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "Example 1",
			arr:      []int{10, 2, 5, 3},
			expected: true,
		},
		{
			name:     "Example 2",
			arr:      []int{3, 1, 7, 11},
			expected: false,
		},
		{
			name:     "Example 3",
			arr:      []int{7, 1, 14, 11},
			expected: true,
		},
		{
			name:     "Edge Case",
			arr:      []int{0, 0},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkIfExist(tt.arr)
			assert.Equal(t, tt.expected, result)
		})
	}
}
