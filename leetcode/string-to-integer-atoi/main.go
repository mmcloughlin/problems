package main

import "log"

const (
	intmin = -(1 << 31)
	intmax = (1 << 31) - 1
)

func isdigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func digitvalue(b byte) int64 {
	return int64(b - '0')
}

func myAtoi(str string) int {
	i, n := 0, len(str)

	// Skip initial spaces.
	for ; i < n && str[i] == ' '; i++ {
	}

	// Consume optional sign character.
	if i == n {
		return 0
	}

	s := int64(1)
	switch str[i] {
	case '+':
		i++
	case '-':
		s = -1
		i++
	}

	// Expect at least one digit.
	if i == n || !isdigit(str[i]) {
		return 0
	}

	var x int64
	for ; i < n && isdigit(str[i]); i++ {
		x = 10*x + digitvalue(str[i])

		// Handle over/underflow.
		if (s * x) < intmin {
			return intmin
		}
		if (s * x) > intmax {
			return intmax
		}
	}

	return int(s * x)
}

func main() {
	cases := []struct {
		Input  string
		Expect int
	}{
		{"42", 42},
		{"     -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"    ", 0},
		{"   +", 0},
		{" -", 0},
		{" +w", 0},
		{"9223372036854775808", 2147483647},
	}
	for _, c := range cases {
		got := myAtoi(c.Input)
		if got != c.Expect {
			log.Fatalf("myAtoi(%q) = %d; expect %d", c.Input, got, c.Expect)
		}
	}
	log.Print("pass")
}
