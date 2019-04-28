package solution

import "testing"

func TestLFUCacheExample(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	asserteq(t, c.Get(1), 1)
	c.Put(3, 3) // evicts key 2
	asserteq(t, c.Get(2), -1)
	asserteq(t, c.Get(3), 3)
	c.Put(4, 4) // evicts key 1.
	asserteq(t, c.Get(1), -1)
	asserteq(t, c.Get(3), 3)
	asserteq(t, c.Get(4), 4)
}

func TestLFUCacheExample2(t *testing.T) {
	// ["LFUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// expect: [null,null,null,null,null,-1,3]
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	asserteq(t, c.Get(1), -1)
	asserteq(t, c.Get(2), 3)
}

func TestLFUCacheEmpty(t *testing.T) {
	c := Constructor(0)
	c.Put(0, 0)
	asserteq(t, c.Get(0), -1)
}

func TestLFUCacheSingle(t *testing.T) {
	c := Constructor(1)
	for n := 1; n <= 10; n++ {
		c.Put(n, n)
		for i := 0; i < n; i++ {
			asserteq(t, c.Get(n), n)
		}
	}
}

func asserteq(t *testing.T, got, expect int) {
	if got != expect {
		t.Errorf("got %d expect %d", got, expect)
	}
}
