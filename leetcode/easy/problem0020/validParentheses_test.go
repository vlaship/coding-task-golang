package problem0020

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValidParentheses_1(t *testing.T) {
	expected := true
	result := isValidParentheses("()")
	require.Equal(t, expected, result)
}

func Test_isValidParentheses_2(t *testing.T) {
	expected := true
	result := isValidParentheses("()[]{}")
	require.Equal(t, expected, result)
}

func Test_isValidParentheses_3(t *testing.T) {
	expected := false
	result := isValidParentheses("(]")
	require.Equal(t, expected, result)
}

func Test_isValidParentheses_4(t *testing.T) {
	expected := false
	result := isValidParentheses("){")
	require.Equal(t, expected, result)
}

func Test_isValidParentheses_5(t *testing.T) {
	expected := false
	result := isValidParentheses("(){}}{")
	require.Equal(t, expected, result)
}
