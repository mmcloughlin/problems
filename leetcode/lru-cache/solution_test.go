package solution

import "testing"

func TestLRUCacheExample(t *testing.T) {
	cache := Constructor(2)
	c := &cache

	c.Put(1, 1)
	c.Put(2, 2)
	get(t, c, 1, 1)

	// evicts 2
	c.Put(3, 3)
	get(t, c, 2, -1)

	// evicts 1
	c.Put(4, 4)
	get(t, c, 1, -1)

	// should contain 3 and 4
	get(t, c, 3, 3)
	get(t, c, 4, 4)

	// change a value
	c.Put(3, 333)
	get(t, c, 3, 333)

	c.Put(5, 5)
	get(t, c, 3, 333)
	get(t, c, 4, -1)
	get(t, c, 5, 5)
}

func get(t *testing.T, c *LRUCache, key, expect int) {
	if got := c.Get(key); got != expect {
		t.Errorf("Get(%v) = %v; expect %v", key, got, expect)
	}
}
