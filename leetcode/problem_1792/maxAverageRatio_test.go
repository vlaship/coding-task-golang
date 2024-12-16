package problem_1792

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxAverageRatio(t *testing.T) {
	tests := []struct {
		name          string
		classes       [][]int
		extraStudents int
		want          float64
	}{
		{
			name:          "Example 1",
			classes:       [][]int{{1, 2}, {3, 5}, {2, 2}},
			extraStudents: 2,
			want:          0.78333,
		},
		{
			name:          "Example 2",
			classes:       [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
			extraStudents: 4,
			want:          0.53485,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAverageRatio(tt.classes, tt.extraStudents)
			// Using a small epsilon for floating point comparison
			require.True(t, math.Abs(got-tt.want) < 1e-5, "got %v, want %v", got, tt.want)
		})
	}
}
