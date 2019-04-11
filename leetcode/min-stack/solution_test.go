package solution

import (
	"math"
	"math/rand"
	"testing"
)

func TestMinStackRepeatValue(t *testing.T) {
	s := Constructor()
	n := 100
	x := 7

	for i := 0; i < n; i++ {
		s.Push(x)
		if s.GetMin() != x {
			t.Fail()
		}
	}

	for i := 0; i < n; i++ {
		if s.GetMin() != x {
			t.Fail()
		}
		s.Pop()
	}
}

func TestMinStackDecreasing(t *testing.T) {
	s := Constructor()
	n := 100
	x := 700

	for i := 0; i < n; i++ {
		s.Push(x)
		if s.GetMin() != x {
			t.Fail()
		}
		x--
	}
}

func TestMinStackIncreasing(t *testing.T) {
	s := Constructor()
	n := 100
	b := 33
	x := b

	for i := 0; i < n; i++ {
		s.Push(x)
		if s.GetMin() != b {
			t.Fail()
		}
		x++
	}
}

func TestMinStackAlternate(t *testing.T) {
	s := Constructor()

	s.Push(2)
	s.Push(1)
	s.Push(2)
	s.Push(1)

	for _, expect := range []int{1, 1, 1, 2} {
		if s.GetMin() != expect {
			t.Fail()
		}
		s.Pop()
	}
}

func TestMinStackRandom(t *testing.T) {
	s := Constructor()
	for i := 0; i < 100000; i++ {
		s.Push(rand.Intn(100))

		for j := 0; j < rand.Intn(1+len(s.stack)/2)-1; j++ {
			s.Pop()
		}

		if s.GetMin() != min(s.stack) {
			t.Fail()
		}
	}
}

func min(xs []int) int {
	m := math.MaxInt64
	for _, x := range xs {
		if x < m {
			m = x
		}
	}
	return m
}
