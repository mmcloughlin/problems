package solution

import "testing"

func TestCharCounts(t *testing.T) {
	s := "bcdacebe"
	expect := "(a,1),(b,2),(c,2),(d,1),(e,2)"
	if expect != CharCounts(s) {
		t.Fail()
	}
}
