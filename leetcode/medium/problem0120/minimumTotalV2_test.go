package problem0120

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTotalV2_1(t *testing.T) {
	expected := 11
	result := minimumTotalV2([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	require.Equal(t, expected, result)
}
func Test_minimumTotalV2_2(t *testing.T) {
	expected := -10
	result := minimumTotalV2([][]int{{-10}})
	require.Equal(t, expected, result)
}

//2(2)
//3(5) 4(6)
//6(11) 5(10,11->10) 9(15)
//4(15) 4(15,14->14) 8(19,23->19) 0(15)
//5(20) 6(21,21->21) 7(21,26->21) 8(31,23->23) 9(24)

func Test_minimumTotalV2_3(t *testing.T) {
	expected := 14
	result := minimumTotalV2([][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}})
	require.Equal(t, expected, result)
}
func Test_minimumTotalV2_4(t *testing.T) {
	expected := 20
	result := minimumTotalV2([][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}, {5, 6, 7, 8, 9}})
	require.Equal(t, expected, result)
}
