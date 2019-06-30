package main

type node struct {
	prev  *node
	next  *node
	key   int
	value int
}

type LRUCache struct {
	kv       map[int]*node
	head     *node
	tail     *node
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		kv:       make(map[int]*node),
		capacity: capacity,
		head:     &node{},
		tail:     &node{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (cache *LRUCache) Get(key int) int {
	if n, ok := cache.kv[key]; ok {
		cache.remove(n)
		cache.add(n)
		return n.value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if n, ok := cache.kv[key]; ok {
		cache.remove(n)
		n = &node{key: key, value: value}
		cache.add(n)

	} else {
		if len(cache.kv) >= cache.capacity {
			cache.remove(cache.tail.prev)
		}
		n = &node{key: key, value: value}
		cache.add(n)
	}
}

func (cache *LRUCache) add(n *node) {
	n.prev = cache.head
	n.next = cache.head.next
	n.prev.next = n
	n.next.prev = n

	cache.kv[n.key] = n
}

func (cache *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil

	delete(cache.kv, n.key)
}

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	cache.Get(4)
}
