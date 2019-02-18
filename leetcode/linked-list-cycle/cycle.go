package cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for {
		slow = next(slow)
		fast = next(next(fast))
		if slow == nil || fast == nil {
			return false
		}
		if slow == fast {
			return true
		}
	}
}

func next(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	return n.Next
}
