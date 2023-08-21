package problem0448

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDisappearedNumbers_1(t *testing.T) {
	expected := []int{5, 6}
	result := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	require.Equal(t, expected, result)
}

func Test_findDisappearedNumbers_2(t *testing.T) {
	expected := []int{2}
	result := findDisappearedNumbers([]int{1, 1})
	require.Equal(t, expected, result)
}
