package problem0217

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate_1(t *testing.T) {
	expected := true
	result := containsDuplicate([]int{1, 2, 3, 1})
	require.Equal(t, expected, result)
}

func Test_containsDuplicate_2(t *testing.T) {
	expected := false
	result := containsDuplicate([]int{1, 2, 3, 4})
	require.Equal(t, expected, result)
}

func Test_containsDuplicate_3(t *testing.T) {
	expected := true
	result := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	require.Equal(t, expected, result)
}
