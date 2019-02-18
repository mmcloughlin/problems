package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := &ListNode{}
	cur := sentinel

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return sentinel.Next
}
