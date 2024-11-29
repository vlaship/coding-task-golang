package problem0412

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "Example 1",
			input:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "Example 2",
			input:    5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:     "Example 3",
			input:    15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fizzBuzz(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
