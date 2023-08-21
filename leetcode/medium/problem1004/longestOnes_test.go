package problem1004

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestOnes_1(t *testing.T) {
	expected := 6
	result := longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	require.Equal(t, expected, result)
}
func Test_longestOnes_2(t *testing.T) {
	expected := 10
	result := longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	require.Equal(t, expected, result)
}
