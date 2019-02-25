package solution

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = step(slow)
		fast = step(step(fast))
		if slow == 1 || fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func step(n int) int {
	s := 0
	for n != 0 {
		d := n % 10
		s += d * d
		n /= 10
	}
	return s
}
