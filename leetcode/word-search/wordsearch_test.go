package wordsearch

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	cases := []struct {
		Word   string
		Expect bool
	}{
		{"A", true},
		{"C", true},
		{"AB", true},
		{"FCS", true},
		{"ABCCED", true},
		{"SEE", true},
		{"ABCB", false},
	}
	for _, c := range cases {
		t.Run(c.Word, func(t *testing.T) {
			got := exist(board, c.Word)
			if got != c.Expect {
				t.Errorf("exist(..., %q) = %v; expect %v", c.Word, got, c.Expect)
			}
		})
	}
}
