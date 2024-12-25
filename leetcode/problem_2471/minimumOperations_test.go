package problem_2471

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func Test_minimumOperations(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected int
	}{
		{
			name:     "Example 1",
			input:    []interface{}{1, 4, 3, 7, 6, 8, 5, nil, nil, nil, nil, 9, nil, 10},
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    []interface{}{1, 3, 2, 7, 6, 5, 4},
			expected: 3,
		},
		{
			name:     "Example 3",
			input:    []interface{}{1, 2, 3, 4, 5, 6},
			expected: 0,
		},
		{
			name:     "Single node",
			input:    []interface{}{1},
			expected: 0,
		},
		{
			name:     "Empty tree",
			input:    []interface{}{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.input)
			result := minimumOperations(root)
			require.Equal(t, tt.expected, result)
		})
	}
}
