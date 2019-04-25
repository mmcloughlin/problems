package solution

const idle = 0

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	// Count tasks.
	ondeck := map[byte]int{}
	total := 0
	for _, task := range tasks {
		ondeck[task]++
		total++
	}

	// Window of the last n intervals.
	window := make([]byte, n)

	// Tasks in cooldown period.
	cool := map[byte]int{}

	// Process tasks.
	intervals := 0
	t := 0
	for total > 0 {
		// Pick a task to run.
		run := pick(ondeck)

		// Task coming off cool-down period?
		if window[t] != idle {
			task := window[t]
			ondeck[task] = cool[task]
			delete(cool, task)
		}

		window[t] = run
		if run != idle {
			if ondeck[run] > 1 {
				cool[run] = ondeck[run] - 1
			}
			delete(ondeck, run)
			total--
		}

		intervals++
		t = (t + 1) % n
	}

	return intervals
}

func pick(tasks map[byte]int) byte {
	picked := byte(idle)
	max := 0
	for task, count := range tasks {
		if count > max {
			picked = task
			max = count
		}
	}
	return picked
}
