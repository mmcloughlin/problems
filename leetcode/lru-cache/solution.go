package solution

const notfound = -1

type item struct {
	key   int
	value int
	prev  *item
	next  *item
}

type LRUCache struct {
	capacity int
	items    map[int]*item
	head     *item
	tail     *item
}

func Constructor(capacity int) LRUCache {
	head, tail := &item{}, &item{}
	tail.next = head
	head.prev = tail
	return LRUCache{
		capacity: capacity,
		items:    map[int]*item{},
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	i, ok := this.items[key]
	if !ok {
		return notfound
	}

	// prune out the item
	i.prev.next = i.next
	i.next.prev = i.prev

	// move to newest position
	prev := this.head.prev
	prev.next = i
	i.prev = prev
	i.next = this.head
	this.head.prev = i

	return i.value
}

func (this *LRUCache) Put(key int, value int) {
	// Is it already there?
	if this.Get(key) != notfound {
		this.items[key].value = value
		return
	}

	// Have we reached capacity? If so remove the LRU.
	if this.Size() == this.capacity {
		lru := this.tail.next
		this.tail.next = lru.next
		lru.next.prev = this.tail
		delete(this.items, lru.key)
	}

	// Add it.
	i := &item{
		key:   key,
		value: value,
		prev:  this.head,
		next:  nil,
	}
	prev := this.head.prev
	prev.next = i
	i.prev = prev
	i.next = this.head
	this.head.prev = i
	this.items[key] = i
}

func (this LRUCache) Size() int {
	return len(this.items)
}
