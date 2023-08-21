package problem0001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSumV2(t *testing.T) {
	expected := [][]int{{3, 3}, {1, 5}, {4, 2}}
	result := twoSumV2([]int{1, 3, 3, 4, 5, 2}, 6)
	require.Equal(t, expected, result)
}
