package problem0121

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit_1(t *testing.T) {
	expected := 5
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	require.Equal(t, expected, result)
}

func Test_maxProfit_2(t *testing.T) {
	expected := 0
	result := maxProfit([]int{7, 6, 4, 3, 1})
	require.Equal(t, expected, result)
}

func Test_maxProfit_3(t *testing.T) {
	expected := 0
	result := maxProfit([]int{1})
	require.Equal(t, expected, result)
}

func Test_maxProfit_4(t *testing.T) {
	expected := 1
	result := maxProfit([]int{1, 2})
	require.Equal(t, expected, result)
}

func Test_maxProfit_5(t *testing.T) {
	expected := 3
	result := maxProfit([]int{1, 4, 2})
	require.Equal(t, expected, result)
}

func Test_maxProfit_6(t *testing.T) {
	expected := 3
	result := maxProfit([]int{2, 1, 4})
	require.Equal(t, expected, result)
}

func Test_maxProfit_7(t *testing.T) {
	expected := 1
	result := maxProfit([]int{2, 1, 2, 0, 1})
	require.Equal(t, expected, result)
}
