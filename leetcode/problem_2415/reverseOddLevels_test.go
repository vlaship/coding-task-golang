package problem_2415

import (
	"reflect"
	"testing"
)

func createTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
			i++
		}
		if i < len(values) {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}

func treeToArray(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func TestReverseOddLevels(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{2, 3, 5, 8, 13, 21, 34},
			expected: []int{2, 5, 3, 8, 13, 21, 34},
		},
		{
			name:     "Example 2",
			input:    []int{7, 13, 11},
			expected: []int{7, 11, 13},
		},
		{
			name:     "Example 3",
			input:    []int{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2},
			expected: []int{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTree(tt.input)
			result := reverseOddLevels(root)
			resultArray := treeToArray(result)

			if !reflect.DeepEqual(resultArray, tt.expected) {
				t.Errorf("got %v, want %v", resultArray, tt.expected)
			}
		})
	}
}
