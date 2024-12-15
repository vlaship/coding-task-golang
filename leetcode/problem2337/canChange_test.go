package problem2337

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canChange(t *testing.T) {
	tests := []struct {
		name   string
		start  string
		target string
		want   bool
	}{
		{
			name:   "Example 1",
			start:  "_L__R__R_",
			target: "L______RR",
			want:   true,
		},
		{
			name:   "Example 2",
			start:  "R_L_",
			target: "__LR",
			want:   false,
		},
		{
			name:   "Example 3",
			start:  "_R",
			target: "R_",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canChange(tt.start, tt.target)
			require.Equal(t, tt.want, got)
		})
	}
}
