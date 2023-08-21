package problem0001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum_1(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum([]int{2, 7, 11, 15}, 9)
	require.Equal(t, expected, result)
}

func Test_twoSum_2(t *testing.T) {
	expected := []int{1, 2}
	result := twoSum([]int{3, 2, 4}, 6)
	require.Equal(t, expected, result)
}

func Test_twoSum_3(t *testing.T) {
	expected := []int{0, 1}
	result := twoSum([]int{3, 3}, 6)
	require.Equal(t, expected, result)
}
