package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nA := length(headA)
	nB := length(headB)

	// Step the longer one to meet the length of the shorter.
	nodeA, nodeB := headA, headB
	if nA > nB {
		nodeA = step(nodeA, nA-nB)
	} else {
		nodeB = step(nodeB, nB-nA)
	}

	// Step in unison until they hit each other.
	for nodeA != nodeB {
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}

	return nodeA
}

func length(head *ListNode) int {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	return n
}

func step(node *ListNode, n int) *ListNode {
	for ; n > 0; n-- {
		node = node.Next
	}
	return node
}
