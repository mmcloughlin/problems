package solution

import "math/rand"

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.clone()
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	s := this.clone()
	for i := 0; i < len(s); i++ {
		j := i + rand.Intn(len(s)-i)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (this *Solution) clone() []int {
	return append([]int{}, this.nums...)
}
