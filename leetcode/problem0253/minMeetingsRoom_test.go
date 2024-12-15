package problem0253

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  2,
		},
		{
			name:      "Example 2",
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMeetingRooms(tt.intervals)
			assert.Equal(t, tt.expected, result)
		})
	}
}
