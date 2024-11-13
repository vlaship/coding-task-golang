package problem0366

/*
366. Find Leaves of Binary Tree

https://leetcode.com/problems/find-leaves-of-binary-tree/

Given the root of a binary tree, collect a tree's nodes as if you were doing this:
Collect all the leaf nodes.
Remove all the leaf nodes.
Repeat until the tree is empty.

Example 1:
Input: root = [1,2,3,4,5]
Output: [[4,5,3],[2],[1]]

Example 2:
Input: root = [1]
Output: [[1]]

Constraints:
The number of nodes in the tree is in the range [1, 100].
-100 <= Node.val <= 100
*/
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}
	level(root, &res)

	return res
}

func level(root *TreeNode, res *[][]int) int {
	if root == nil {
		return -1
	}
	lev := maxLevel(level(root.left, res), level(root.right, res)) + 1

	if len(*res) <= lev {
		*res = append(*res, make([]int, 0))
	}
	(*res)[lev] = append((*res)[lev], root.val)

	return lev
}

func maxLevel(l, r int) int {
	if l > r {
		return l
	}
	return r
}
