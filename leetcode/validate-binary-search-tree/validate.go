package bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTBetween(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTBetween(root *TreeNode, lo, hi int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lo || root.Val >= hi {
		return false
	}
	return isValidBSTBetween(root.Left, lo, root.Val) && isValidBSTBetween(root.Right, root.Val, hi)
}
