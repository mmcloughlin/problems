package solution

import "strconv"

func fizzBuzz(n int) []string {
	lines := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		lines = append(lines, write(i))
	}
	return lines
}

func write(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(i)
}
