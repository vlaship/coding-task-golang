package medium

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
