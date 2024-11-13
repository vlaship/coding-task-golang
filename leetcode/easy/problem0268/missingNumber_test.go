package problem0268

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example1", []int{3, 0, 1}, 2},
		{"example2", []int{0, 1}, 2},
		{"example3", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := missingNumber(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
