package problem2593

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindScore(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{"Example 1", []int{2, 1, 3, 4, 5, 2}, 7},
		{"Example 2", []int{2, 3, 5, 1, 3, 2}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findScore(tt.nums)
			require.Equal(t, tt.want, got, "findScore(%v) = %d; want %d", tt.nums, got, tt.want)
		})
	}
}
