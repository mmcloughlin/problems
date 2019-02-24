package solution

import (
	"testing"
	"testing/quick"
)

func TestGetSum(t *testing.T) {
	expect := func(a, b int) int { return a + b }
	if err := quick.CheckEqual(getSum, expect, nil); err != nil {
		t.Fatal(err)
	}
}
