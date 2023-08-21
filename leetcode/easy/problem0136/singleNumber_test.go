package problem0136

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber_1(t *testing.T) {
	expected := 1
	result := singleNumber([]int{2, 2, 1})
	require.Equal(t, expected, result)
}

func Test_singleNumber_2(t *testing.T) {
	expected := 4
	result := singleNumber([]int{4, 1, 2, 1, 2})
	require.Equal(t, expected, result)
}

func Test_singleNumber_3(t *testing.T) {
	expected := 1
	result := singleNumber([]int{1})
	require.Equal(t, expected, result)
}
