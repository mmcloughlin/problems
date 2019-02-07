package oddeven

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(node *ListNode) *ListNode {
	heads := [2]*ListNode{{}, {}}
	cur := [2]*ListNode{heads[0], heads[1]}

	i := 0
	for node != nil {
		cur[i].Next = node
		cur[i] = node
		node = node.Next
		i ^= 1
	}

	cur[0].Next = heads[1].Next
	cur[1].Next = nil

	return heads[0].Next
}
