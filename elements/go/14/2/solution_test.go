package solution

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	for trial := 0; trial < 100; trial++ {
		m := 10 + rand.Intn(100)
		n := 10 + rand.Intn(100)

		// Generate underlying arrays.
		x := randints(m, 10)
		y := randints(n, 10)

		// Prepare inputs.
		a := make([]int, m+n)
		b := make([]int, n)
		copy(a, x)
		copy(b, y)

		// Execute.
		merge(a, b, m)

		// Check.
		expect := append([]int{}, x...)
		expect = append(expect, y...)
		sort.Ints(expect)

		if !reflect.DeepEqual(expect, a) {
			t.Fail()
		}
	}
}

func randints(n, delta int) []int {
	nums := make([]int, n)
	num := 0
	for i := 0; i < n; i++ {
		num += rand.Intn(delta)
		nums[i] = num
	}
	return nums
}
