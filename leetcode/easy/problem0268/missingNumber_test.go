package problem0268

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingNumber_1(t *testing.T) {
	expected := 2
	result := missingNumber([]int{3, 0, 1})
	require.Equal(t, expected, result)
}

func Test_missingNumber_2(t *testing.T) {
	expected := 2
	result := missingNumber([]int{0, 1})
	require.Equal(t, expected, result)
}

func Test_missingNumber_3(t *testing.T) {
	expected := 8
	result := missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	require.Equal(t, expected, result)
}
