package solution

import "testing"

func TestRandomizedSet(t *testing.T) {
	randset := Constructor()
	s := &randset

	assert(t, s.Insert(1))
	assert(t, !s.Remove(2))
	assert(t, s.GetRandom() == 1)
	assert(t, s.Insert(2))
	x := s.GetRandom()
	assert(t, x == 1 || x == 2)
	assert(t, s.Remove(1))
	assert(t, !s.Insert(2))
	assert(t, s.GetRandom() == 2)
}

func assert(t *testing.T, v bool) {
	if !v {
		t.FailNow()
	}
}
