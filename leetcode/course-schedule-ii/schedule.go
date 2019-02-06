package schedule

func findOrder(numCourses int, prerequisites [][]int) []int {
	requires := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		requires[p[0]] = append(requires[p[0]], p[1])
	}

	const (
		white int = iota // course is planned
		grey             // completing prerequisites
		black            // complete
	)
	state := make(map[int]int, numCourses)

	stack := []int{}
	for c := 0; c < numCourses; c++ {
		state[c] = white
		stack = append(stack, c)
	}

	order := []int{}

	for len(stack) > 0 {
		top := len(stack) - 1
		c := stack[top]
		stack = stack[:top]

		switch state[c] {
		case white:
			state[c] = grey
			stack = append(stack, c)
			for _, r := range requires[c] {
				if state[r] == grey {
					return nil
				}
				stack = append(stack, r)
			}
		case grey:
			state[c] = black
			order = append(order, c)
		case black:
			// nothing to be done
		}
	}

	return order
}
