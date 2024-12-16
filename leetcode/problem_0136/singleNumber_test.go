package problem_0136

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example1", []int{2, 2, 1}, 1},
		{"example2", []int{4, 1, 2, 1, 2}, 4},
		{"example3", []int{1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
