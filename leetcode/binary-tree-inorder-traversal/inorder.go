package inorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	type state struct {
		node        *TreeNode
		visitedleft bool
	}
	values := []int{}
	stack := []state{{root, false}}
	for len(stack) > 0 {
		// Pop the stack.
		top := len(stack) - 1
		s := stack[top]
		stack = stack[:top]

		if s.visitedleft {
			values = append(values, s.node.Val)
			stack = append(stack, state{s.node.Right, false})
			continue
		}

		if s.node == nil {
			continue
		}

		stack = append(stack, state{s.node, true}, state{s.node.Left, false})
	}
	return values
}

/*
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	values := inorderTraversal(root.Left)
	values = append(values, root.Val)
	values = append(values, inorderTraversal(root.Right)...)
	return values
}
*/
