package main

import "log"

func twoSum(nums []int, target int) []int {
	idx := map[int][]int{}
	for i, n := range nums {
		idx[n] = append(idx[n], i)
	}

	for n, in := range idx {
		m := target - n
		if m == n {
			if len(in) >= 2 {
				return in[:2]
			}
			continue
		}
		im, found := idx[m]
		if !found {
			continue
		}
		return []int{in[0], im[0]}
	}
	return nil
}

func main() {
	cases := []struct {
		Nums   []int
		Target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{2, 9, 100, 2}, 4},
		{[]int{1, 2, 3}, 4},
	}

	for _, c := range cases {
		idx := twoSum(c.Nums, c.Target)
		got := c.Nums[idx[0]] + c.Nums[idx[1]]
		if got != c.Target {
			log.Fatal("fail")
		}
	}

	log.Print("pass")
}
