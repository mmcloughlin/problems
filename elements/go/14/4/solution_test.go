package solution

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	x := []int{
		1, 1, 1, 1, 1,
		2, 2, 2,
		3, 3, 3, 3,
		4, 4,
		5, 5, 5, 5, 5, 5, 5, 5,
		6, 6,
		7, 7, 7, 7, 7,
		8, 8, 8,
	}
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}

	input := append([]int{}, x...)
	got := Unique(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Unique(%v) = %v; expect %v", x, got, expect)
	}
}
