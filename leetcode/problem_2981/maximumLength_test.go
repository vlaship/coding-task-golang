package problem_2981

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumLength(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "aaaa",
			want: 2,
		},
		{
			name: "Example 2",
			s:    "abcdef",
			want: -1,
		},
		{
			name: "Example 3",
			s:    "abcaba",
			want: 1,
		},
		{
			name: "Example 4",
			s:    "abcccccdddd",
			want: 3,
		},
		{
			name: "Example 5",
			s:    "ttttt",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumLength(tt.s)
			require.Equal(t, tt.want, got)
		})
	}
}
