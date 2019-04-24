package solution

func dailyTemperatures(T []int) []int {
	n := len(T)
	wait := make([]int, n)
	warmer := []int{}
	for i := n - 1; i >= 0; i-- {
		// Update the stack of warmer days in future.
		for len(warmer) > 0 && T[peek(warmer)] <= T[i] {
			warmer = pop(warmer)
		}

		// Set the number of days we have to wait.
		if len(warmer) > 0 {
			wait[i] = peek(warmer) - i
		}

		warmer = append(warmer, i)
	}

	return wait
}

func peek(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack []int) []int {
	return stack[:len(stack)-1]
}
