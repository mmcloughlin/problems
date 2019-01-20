package remnth

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	// Ahead will be n steps ahead of behind.
	behind, ahead := head, head

	for i := 0; i < n; i++ {
		ahead = ahead.Next
	}

	if ahead == nil {
		return head.Next
	}

	// Step until ahead gets to the end.
	for ahead.Next != nil {
		ahead = ahead.Next
		behind = behind.Next
	}

	// Remove
	behind.Next = behind.Next.Next

	return head
}
