package problem_0006

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		numRows  int
		expected string
	}{
		{
			name:     "Example 1",
			input:    "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Example 2",
			input:    "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Example 3",
			input:    "A",
			numRows:  1,
			expected: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convert(tt.input, tt.numRows)
			require.Equal(t, tt.expected, result)
		})
	}
}
