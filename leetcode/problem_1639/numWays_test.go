package problem_1639

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumWays(t *testing.T) {
	tests := []struct {
		words    []string
		target   string
		expected int
	}{
		{
			words:    []string{"acca", "bbbb", "caca"},
			target:   "aba",
			expected: 6,
		},
		{
			words:    []string{"abba", "baab"},
			target:   "bab",
			expected: 4,
		},
	}

	for _, test := range tests {
		got := numWays(test.words, test.target)
		require.Equal(t, test.expected, got)
	}
}
