package problem_0070

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"example1", 2, 2},
		{"example2", 3, 3},
		{"example3", 4, 5},
		{"example4", 5, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := climbStairs(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
