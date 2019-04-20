package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	return equal(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equal(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && equal(s.Left, t.Left) && equal(s.Right, t.Right)
}
