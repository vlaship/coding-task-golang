package problem_0366

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLeaves(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "example1",
			root: func() *TreeNode {
				n3 := &TreeNode{val: 3, left: nil, right: nil}
				n4 := &TreeNode{val: 4, left: nil, right: nil}
				n5 := &TreeNode{val: 5, left: nil, right: nil}
				n2 := &TreeNode{val: 2, left: n4, right: n5}
				n1 := &TreeNode{val: 1, left: n2, right: n3}
				return n1
			}(),
			expected: [][]int{{4, 5, 3}, {2}, {1}},
		},
		{
			name: "example2",
			root: func() *TreeNode {
				n1 := &TreeNode{val: 1, left: nil, right: nil}
				return n1
			}(),
			expected: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLeaves(tt.root)
			require.Equal(t, tt.expected, result)
		})
	}
}
