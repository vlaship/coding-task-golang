package problem0001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_problem0001(t *testing.T) {
	expected := []int{0, 1, 9, 16, 16, 100}
	result := problem001([]int{-4, -1, 0, 3, 4, 10})
	require.Equal(t, expected, result)
}
