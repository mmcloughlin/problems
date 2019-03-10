package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	mindiff, _, _ := getMinimumDifferenceMinMax(root)
	return mindiff
}

func getMinimumDifferenceMinMax(root *TreeNode) (int, int, int) {
	if root == nil {
		panic("unreachable")
	}

	mindiff := math.MaxInt64
	mn := root.Val
	mx := root.Val

	if root.Left != nil {
		lmindiff, lmin, lmax := getMinimumDifferenceMinMax(root.Left)
		mn = lmin
		mindiff = min(mindiff, lmindiff)
		mindiff = min(mindiff, abs(root.Val-lmax))
	}

	if root.Right != nil {
		rmindiff, rmin, rmax := getMinimumDifferenceMinMax(root.Right)
		mx = rmax
		mindiff = min(mindiff, rmindiff)
		mindiff = min(mindiff, abs(root.Val-rmin))
	}

	return mindiff, mn, mx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
