package problem0020

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"example1", "()", true},
		{"example2", "()[]{}", true},
		{"example3", "(]", false},
		{"example4", "){", false},
		{"example5", "(){}}{", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidParentheses(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
