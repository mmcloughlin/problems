package hamming

import (
	"math/bits"
	"testing"
	"testing/quick"
)

func TestHamming(t *testing.T) {
	expect := func(x, y int) int {
		return bits.OnesCount64(uint64(x) ^ uint64(y))
	}
	if err := quick.CheckEqual(hammingDistance, expect, nil); err != nil {
		t.Fatal(err)
	}
}
