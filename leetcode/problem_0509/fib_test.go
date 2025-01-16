package problem_0509

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{30, 832040},
	}

	for _, tc := range testCases {
		t.Run("Fibonacci Test", func(t *testing.T) {
			result := fib(tc.input)
			require.Equal(t, tc.expected, result, "For input %d, expected %d but got %d", tc.input, tc.expected, result)
		})
	}
}
