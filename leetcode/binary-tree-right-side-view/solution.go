package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	vl := rightSideView(root.Left)
	vr := rightSideView(root.Right)

	v := []int{root.Val}
	v = append(v, vr...)
	if len(vl) > len(vr) {
		v = append(v, vl[len(vr):]...)
	}

	return v
}
