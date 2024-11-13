package problem0448

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"example1", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{"example2", []int{1, 1}, []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDisappearedNumbers(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
