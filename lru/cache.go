package lru

import "container/list"

type KeyPair struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int                   //defines a cache object of the specified capacity.
	list     *list.List            //DoublyLinkedList for backing the cache value.
	elements map[int]*list.Element //Map to store list pointer of cache mapped to key
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.elements[key]; ok {
		value := node.Value.(*list.Element).Value.(KeyPair).value
		cache.list.MoveToFront(node)
		return value
	}
	return -1
}

//Put: Inserts the key,value pair in LRUCache.
//If list capacity is full, entry at the last index of the list is deleted before insertion.
func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.elements[key]; ok {
		cache.list.MoveToFront(node)
		node.Value.(*list.Element).Value = KeyPair{key: key, value: value}
	} else {
		if cache.list.Len() == cache.capacity {
			idx := cache.list.Back().Value.(*list.Element).Value.(KeyPair).key
			delete(cache.elements, idx)
			cache.list.Remove(cache.list.Back())
		}
	}

	node := &list.Element{
		Value: KeyPair{
			key:   key,
			value: value,
		},
	}

	pointer := cache.list.PushFront(node)
	cache.elements[key] = pointer
}

