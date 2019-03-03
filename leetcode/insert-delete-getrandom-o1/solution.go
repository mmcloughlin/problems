package solution

import "math/rand"

type RandomizedSet struct {
	items []int
	index map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		items: []int{},
		index: map[int]int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, found := this.index[val]; found {
		return false
	}
	n := len(this.items)
	this.items = append(this.items, val)
	this.index[val] = n
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	i, found := this.index[val]
	if !found {
		return false
	}
	n := len(this.items)
	this.items[i] = this.items[n-1]
	this.index[this.items[n-1]] = i
	delete(this.index, val)
	this.items = this.items[:n-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n := len(this.items)
	return this.items[rand.Intn(n)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
