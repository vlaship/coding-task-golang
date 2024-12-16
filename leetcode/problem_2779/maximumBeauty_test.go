package problem_2779

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumBeauty(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example1",
			nums: []int{4, 6, 1, 2},
			k:    2,
			want: 3,
		},
		{
			name: "example2",
			nums: []int{1, 1, 1, 1},
			k:    10,
			want: 4,
		},
		{
			name: "example3",
			nums: []int{1, 3, 5, 7},
			k:    1,
			want: 1,
		},
		{
			name: "example4",
			nums: []int{1, 2, 2, 3},
			k:    1,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumBeauty(tt.nums, tt.k)
			require.Equal(t, tt.want, got)
		})
	}
}
