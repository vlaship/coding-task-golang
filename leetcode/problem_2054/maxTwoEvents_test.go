package problem_2054

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoBestNonOverlappingEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name: "Example 1",
			events: [][]int{
				{1, 3, 2},
				{4, 5, 2},
				{2, 4, 3},
			},
			want: 4,
		},
		{
			name: "Example 2",
			events: [][]int{
				{1, 3, 2},
				{4, 5, 2},
				{1, 5, 5},
			},
			want: 5,
		},
		{
			name: "Example 3",
			events: [][]int{
				{1, 5, 3},
				{1, 5, 1},
				{6, 6, 5},
			},
			want: 8,
		},
		{
			name: "Example 4",
			events: [][]int{
				{10, 83, 53},
				{63, 87, 45},
				{97, 100, 32},
				{51, 61, 16},
			},
			want: 85,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTwoEvents(tt.events)
			require.Equal(t, tt.want, got)
		})
	}
}
