package problem_1475

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_finalPrices(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   []int
	}{
		{
			name:   "Example 1",
			prices: []int{8, 4, 6, 2, 3},
			want:   []int{4, 2, 4, 2, 3},
		},
		{
			name:   "Example 2",
			prices: []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Example 3",
			prices: []int{10, 1, 1, 6},
			want:   []int{9, 0, 1, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finalPrices(tt.prices)
			require.Equal(t, tt.want, got)
		})
	}
}
