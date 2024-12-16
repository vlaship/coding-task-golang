package problem_0003

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Example 1", "abcabcbb", 3},
		{"Example 2", "bbbbb", 1},
		{"Example 3", "pwwkew", 3},
		{"Empty String", "", 0},
		{"One character", "a", 1},
		{"All Unique Characters", "abcdef", 6},
		{"Repeated Characters", "aabbcc", 2},
		{"Mixed Characters", "aabacbebebe", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
