package problem0205

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"example1", "egg", "add", true},
		{"example2", "foo", "bar", false},
		{"example3", "paper", "title", true},
		{"example4", "bar", "foo", false},
		{"example5", "", "", true},
		{"example6", "a", "a", true},
		{"example7", "ab", "aa", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isIsomorphic(tt.s, tt.t)
			require.Equal(t, tt.expected, result)
		})
	}
}
