package kthsmallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	v, _ := kthSmallestSize(root, k)
	return v
}

func kthSmallestSize(root *TreeNode, k int) (int, int) {
	if root == nil {
		return 0, 0
	}

	lv, ln := kthSmallestSize(root.Left, k)
	if ln >= k {
		return lv, ln
	}

	if k == ln+1 {
		return root.Val, ln + 1
	}

	rv, rn := kthSmallestSize(root.Right, k-ln-1)

	return rv, ln + 1 + rn
}
