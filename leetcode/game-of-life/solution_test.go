package solution

import (
	"reflect"
	"testing"
)

func TestGameOfLifeExample(t *testing.T) {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	expect := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}

	gameOfLife(board)

	if !reflect.DeepEqual(expect, board) {
		t.Fail()
	}
}
