package main

import "log"

func reverse(x int) int {
	// Determine sign.
	s := 1
	if x < 0 {
		s = -1
		x = -x
	}

	// Compute reversed value.
	r := 0
	for x != 0 {
		r = 10*r + (x % 10)
		x /= 10
	}

	// Handle overflow.
	if r > (1<<31)-1 {
		return 0
	}

	return s * r
}

func main() {
	cases := []struct {
		Input  int
		Expect int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{5, 5},
	}
	for _, c := range cases {
		got := reverse(c.Input)
		if got != c.Expect {
			log.Fatalf("reverse(%d) = %d; expect %d", c.Input, got, c.Expect)
		}
	}
	log.Print("pass")
}
