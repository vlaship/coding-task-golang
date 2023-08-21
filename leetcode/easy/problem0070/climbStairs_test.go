package problem0070

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbStairs_1(t *testing.T) {
	expected := 2
	result := climbStairs(2)
	require.Equal(t, expected, result)
}

func Test_climbStairs_2(t *testing.T) {
	expected := 3
	result := climbStairs(3)
	require.Equal(t, expected, result)
}

/*
Input: n = 4
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 + 1 + 1 + 1
2. 2 + 1 + 1
3. 1 + 2 + 1
4. 1 + 1 + 2
5. 2 + 2
*/
func Test_climbStairs_3(t *testing.T) {
	expected := 5
	result := climbStairs(4)
	require.Equal(t, expected, result)
}

func Test_climbStairs_4(t *testing.T) {
	expected := 8
	result := climbStairs(5)
	require.Equal(t, expected, result)
}
