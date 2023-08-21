package problem0344

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseString(t *testing.T) {
	expected := "!dlroW olleH"
	result := reverseString("Hello World!")
	require.Equal(t, expected, result)
}
