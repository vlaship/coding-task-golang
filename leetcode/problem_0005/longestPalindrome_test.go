package problem_0005

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"example1", "babad", "bab"},
		{"example2", "cbbd", "bb"},
		{"example3", "aaaa", "aaaa"},
		{"example4", "aacabdkacaa", "aca"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestPalindrome(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
