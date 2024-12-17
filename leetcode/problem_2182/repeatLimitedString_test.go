package problem_2182

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatLimitedString(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		repeatLimit int
		want        string
	}{
		{
			name:        "Example 1",
			s:           "cczazcc",
			repeatLimit: 3,
			want:        "zzcccac",
		},
		{
			name:        "Example 2",
			s:           "aababab",
			repeatLimit: 2,
			want:        "bbabaa",
		},
		{
			name:        "Single character",
			s:           "aaaa",
			repeatLimit: 2,
			want:        "aa",
		},
		{
			name:        "All different characters",
			s:           "abcdef",
			repeatLimit: 2,
			want:        "fedcba",
		},
		{
			name:        "Repeat limit 1",
			s:           "zzzaaa",
			repeatLimit: 1,
			want:        "zazaza",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repeatLimitedString(tt.s, tt.repeatLimit)
			require.Equal(t, tt.want, got)
		})
	}
}
