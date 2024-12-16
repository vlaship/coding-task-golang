package problem_2558

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTakeGifts(t *testing.T) {
	tests := []struct {
		name  string
		gifts []int
		k     int
		want  int64
	}{
		{"Example 1", []int{25, 64, 9, 4, 100}, 4, 29},
		{"Example 2", []int{1, 1, 1, 1}, 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pickGifts(tt.gifts, tt.k)
			require.Equal(t, tt.want, got, "takeGifts(%v, %d) = %d; want %d", tt.gifts, tt.k, got, tt.want)
		})
	}
}
