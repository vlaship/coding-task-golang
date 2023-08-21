package problem0005

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome_1(t *testing.T) {
	expected := "bab"
	result := longestPalindrome("babad")
	require.Equal(t, expected, result)
}

func Test_longestPalindrome_2(t *testing.T) {
	expected := "bb"
	result := longestPalindrome("cbbd")
	require.Equal(t, expected, result)
}

func Test_longestPalindrome_3(t *testing.T) {
	expected := "aaaa"
	result := longestPalindrome("aaaa")
	require.Equal(t, expected, result)
}

func Test_longestPalindrome_4(t *testing.T) {
	expected := "aca"
	result := longestPalindrome("aacabdkacaa")
	require.Equal(t, expected, result)
}
