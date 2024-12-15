package problem2825

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeStringSubsequence(t *testing.T) {
	testCases := []struct {
		name   string
		str1   string
		str2   string
		result bool
	}{
		{
			name:   "Example 1",
			str1:   "abc",
			str2:   "ad",
			result: true,
		},
		{
			name:   "Example 2",
			str1:   "zc",
			str2:   "ad",
			result: true,
		},
		{
			name:   "Example 3",
			str1:   "ab",
			str2:   "d",
			result: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canMakeSubsequence(tc.str1, tc.str2)
			require.Equal(t, tc.result, result)
		})
	}
}
