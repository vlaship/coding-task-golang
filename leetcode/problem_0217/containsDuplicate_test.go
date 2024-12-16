package problem_0217

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"example1", []int{1, 2, 3, 1}, true},
		{"example2", []int{1, 2, 3, 4}, false},
		{"example3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			require.Equal(t, tt.expected, result)
		})
	}
}
