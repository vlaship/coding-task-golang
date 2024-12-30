package problem_2466

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountGoodStrings(t *testing.T) {
	tests := []struct {
		name     string
		low      int
		high     int
		zero     int
		one      int
		expected int
	}{
		{
			name:     "Example 1: low=3, high=3, zero=1, one=1",
			low:      3,
			high:     3,
			zero:     1,
			one:      1,
			expected: 8,
		},
		{
			name:     "Example 2: low=2, high=3, zero=1, one=2",
			low:      2,
			high:     3,
			zero:     1,
			one:      2,
			expected: 5,
		},
		{
			name:     "Single length string",
			low:      1,
			high:     1,
			zero:     1,
			one:      1,
			expected: 2,
		},
		{
			name:     "Larger numbers",
			low:      5,
			high:     10,
			zero:     2,
			one:      1,
			expected: 220,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countGoodStrings(tt.low, tt.high, tt.zero, tt.one)
			require.Equal(t, tt.expected, result)
		})
	}
}
