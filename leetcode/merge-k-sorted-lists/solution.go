package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// Handle base cases.
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	// Recurse on the two halfs.
	m := len(lists) / 2
	return merge(
		mergeKLists(lists[:m]),
		mergeKLists(lists[m:]),
	)
}

func merge(a, b *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}

	return head.Next
}
