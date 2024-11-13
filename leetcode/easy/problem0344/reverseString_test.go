package problem0344

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"example1", "Hello World!", "!dlroW olleH"},
		{"example2", "abcd", "dcba"},
		{"example3", "A man, a plan, a canal, Panama", "amanaP ,lanac a ,nalp a ,nam A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseString(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
