package cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// Tortoise and the hare.
	slow, fast := head, head
	steps := 0
	for {
		slow = next(slow)
		fast = next(next(fast))
		steps++

		if slow == nil || fast == nil {
			return nil
		}

		if slow == fast {
			break
		}
	}

	// Now send one ahead by steps (which must be a multiple of header length).
	a, b := head, head
	for s := 0; s < steps; s++ {
		a = a.Next
	}

	for a != b {
		a = a.Next
		b = b.Next
	}

	return a
}

func next(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	return n.Next
}
