package solution

import (
	"reflect"
	"testing"
)

func TestGetSkyline(t *testing.T) {
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	expect := [][]int{
		{2, 10},
		{3, 15},
		{7, 12},
		{12, 0},
		{15, 10},
		{20, 8},
		{24, 0},
	}
	got := getSkyline(buildings)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("getSkyline(%v) = %v; expect %v", buildings, got, expect)
	}
}
