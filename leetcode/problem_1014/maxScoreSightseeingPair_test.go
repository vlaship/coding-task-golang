package problem_1014

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxScoreSightseeingPair(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "example 1",
			values: []int{8, 1, 5, 2, 6},
			want:   11,
		},
		{
			name:   "example 2",
			values: []int{1, 2},
			want:   2,
		},
		{
			name:   "all same values",
			values: []int{1, 1, 1, 1},
			want:   1,
		},
		{
			name:   "decreasing values",
			values: []int{5, 4, 3, 2, 1},
			want:   8, // 5 + 4 + 0 - 1 = 8
		},
		{
			name:   "increasing values",
			values: []int{1, 2, 3, 4, 5},
			want:   8, // 4 + 5 + 3 - 4 = 8
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScoreSightseeingPair(tt.values)
			require.Equal(t, tt.want, got)
		})
	}
}
