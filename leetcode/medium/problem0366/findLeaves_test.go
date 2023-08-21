package problem0366

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLeaves_1(t *testing.T) {
	expected := [][]int{{4, 5, 3}, {2}, {1}}
	n3 := &TreeNode{val: 3, left: nil, right: nil}
	n4 := &TreeNode{val: 4, left: nil, right: nil}
	n5 := &TreeNode{val: 5, left: nil, right: nil}
	n2 := &TreeNode{val: 2, left: n4, right: n5}
	n1 := &TreeNode{val: 1, left: n2, right: n3}
	result := findLeaves(n1)
	require.Equal(t, expected, result)
}
