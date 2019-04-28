package solution

import (
	"fmt"
	"math"
)

// notfound is returned when the key isn't present in the cache.
const notfound = -1

// item represents an item in the cache.
type item struct {
	key  int
	val  int
	freq int
	prev *item
	next *item
	buck *bucket
}

// bucket contains items of the same frequency.
type bucket struct {
	freq  int
	count int
	head  *item
	tail  *item
	prev  *bucket
	next  *bucket
}

func newbucket(freq int) *bucket {
	head, tail := &item{}, &item{}
	head.next = tail
	tail.prev = head
	return &bucket{
		freq:  freq,
		count: 0,
		head:  head,
		tail:  tail,
	}
}

// LFUCache caches a given number of items. At capacity, items are evicted in
// order of least frequently accessed (on ties the least recently accessed is
// chosen).
type LFUCache struct {
	capacity int
	items    map[int]*item

	// sentinel frequency buckets
	head *bucket
	tail *bucket
}

func Constructor(capacity int) LFUCache {
	head := &bucket{freq: math.MinInt64}
	tail := &bucket{freq: math.MaxInt64}
	head.next = tail
	tail.prev = head
	return LFUCache{
		capacity: capacity,
		items:    map[int]*item{},
		head:     head,
		tail:     tail,
	}
}

func (this *LFUCache) Get(key int) int {
	// Fetch the item.
	i, ok := this.items[key]
	if !ok {
		return notfound
	}

	// Delete from its deque.
	i.prev.next = i.next
	i.next.prev = i.prev

	// Add frequency bucket if necessary.
	if i.buck.next.freq != i.freq+1 {
		b := newbucket(i.freq + 1)
		b.prev = i.buck
		b.next = i.buck.next
		b.next.prev = b
		b.prev.next = b
	}

	// Move item to the next frequency bucket.
	i.buck = i.buck.next
	b := i.buck
	i.next = b.tail
	i.prev = b.tail.prev
	b.tail.prev.next = i
	b.tail.prev = i
	b.count++

	// Clean up the previous bucket if its empty.
	b = i.buck.prev
	b.count--
	if b.count == 0 {
		b.prev.next = b.next
		b.next.prev = b.prev
	}

	i.freq++
	return i.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	// Check if its there already, if so update value.
	if i, ok := this.items[key]; ok {
		i.val = value
		this.Get(key)
		return
	}

	// Check if we reached capacity. If so evict.
	if len(this.items) == this.capacity {
		lfb := this.head.next
		lfi := lfb.head.next
		lfi.prev.next = lfi.next
		lfi.next.prev = lfi.prev
		lfb.count--
		if lfb.count == 0 {
			lfb.prev.next = lfb.next
			lfb.next.prev = lfb.prev
		}
		delete(this.items, lfi.key)
	}

	// Ensure we have a bucket for frequency 1.
	if this.head.next.freq != 1 {
		b := newbucket(1)
		b.prev = this.head
		b.next = this.head.next
		b.next.prev = b
		b.prev.next = b
	}

	// Add the item.
	b := this.head.next
	i := &item{key: key, val: value, freq: 1, buck: b}
	i.next = b.tail
	i.prev = b.tail.prev
	b.tail.prev.next = i
	b.tail.prev = i
	b.count++

	this.items[key] = i
}

func (this *LFUCache) Dump() {
	fmt.Println("items")
	for k, i := range this.items {
		fmt.Printf("\t%d -> %d freq=%d\n", k, i.val, i.freq)
	}

	fmt.Println("buckets")
	for b := this.head; b != nil; b = b.next {
		fmt.Printf("\tcount=%d freq=%d\n", b.count, b.freq)
	}
}
