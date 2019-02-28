package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	var pre *ListNode
	for node.Next != nil {
		node.Val = node.Next.Val
		pre, node = node, node.Next
	}
	pre.Next = nil
}
