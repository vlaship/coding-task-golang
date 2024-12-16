package problem_1455

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPrefixOfWord(t *testing.T) {
	tests := []struct {
		name       string
		sentence   string
		searchWord string
		expected   int
	}{
		{
			name:       "Example 1",
			sentence:   "i love eating burger",
			searchWord: "burg",
			expected:   4,
		},
		{
			name:       "Example 2",
			sentence:   "this problem is an easy problem",
			searchWord: "pro",
			expected:   2,
		},
		{
			name:       "Example 3",
			sentence:   "i am tired",
			searchWord: "you",
			expected:   -1,
		},
		{
			name:       "No match",
			sentence:   "hello world",
			searchWord: "test",
			expected:   -1,
		},
		{
			name:       "Match first word",
			sentence:   "test the function",
			searchWord: "test",
			expected:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPrefixOfWord(tt.sentence, tt.searchWord)
			require.Equal(t, tt.expected, result)
		})
	}
}
