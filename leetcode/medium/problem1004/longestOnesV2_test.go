package problem1004

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestOnesV2_1(t *testing.T) {
	expected := 6
	result := longestOnesV2([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
	require.Equal(t, expected, result)
}
func Test_longestOnesV2_2(t *testing.T) {
	expected := 10
	result := longestOnesV2([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)
	require.Equal(t, expected, result)
}
