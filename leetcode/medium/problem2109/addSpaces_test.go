package problem2109

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSpaces(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		spaces   []int
		expected string
	}{
		{
			name:     "Example 1",
			s:        "LeetcodeHelpsMeLearn",
			spaces:   []int{8, 13, 15},
			expected: "Leetcode Helps Me Learn",
		},
		{
			name:     "Example 2",
			s:        "icodeinpython",
			spaces:   []int{1, 5, 7, 9},
			expected: "i code in py thon",
		},
		{
			name:     "Example 3",
			s:        "spacing",
			spaces:   []int{0, 1, 2, 3, 4, 5, 6},
			expected: " s p a c i n g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addSpaces(tt.s, tt.spaces)
			require.Equal(t, tt.expected, result)
		})
	}
}
