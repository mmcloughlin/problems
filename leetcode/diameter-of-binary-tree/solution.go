package solution

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	d, _ := diameterAndDepth(root)
	return d
}

func diameterAndDepth(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	diaml, depl := diameterAndDepth(root.Left)
	diamr, depr := diameterAndDepth(root.Right)
	return max(diaml, diamr, depl+depr), 1 + max(depl, depr)
}

func max(nums ...int) int {
	m := math.MinInt64
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
