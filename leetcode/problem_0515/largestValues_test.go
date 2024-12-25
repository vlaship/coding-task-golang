package problem_0515

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestValues(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			expected: []int{1, 3, 9},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: []int{1, 3},
		},
		{
			name:     "Empty Tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single Node",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestValues(tt.root)
			require.Equal(t, tt.expected, result)
		})
	}
}
