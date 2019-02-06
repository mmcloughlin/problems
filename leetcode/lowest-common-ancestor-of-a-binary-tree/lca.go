package lca

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	a, _ := searchCommonAncestor(root, p, q)
	return a
}

func searchCommonAncestor(root, p, q *TreeNode) (*TreeNode, uint8) {
	if root == nil {
		return nil, 0
	}

	al, ml := searchCommonAncestor(root.Left, p, q)
	if ml == 3 {
		return al, 3
	}

	ar, mr := searchCommonAncestor(root.Right, p, q)
	if mr == 3 {
		return ar, 3
	}

	m := ml | mr
	if root.Val == p.Val {
		m |= 1
	}
	if root.Val == q.Val {
		m |= 2
	}

	if m == 3 {
		return root, 3
	}
	return nil, m
}
