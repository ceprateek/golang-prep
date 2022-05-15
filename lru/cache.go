package lru

import (
	"container/list"
	"fmt"
)

type Cache struct {
	capacity int
	list     *list.List
	elements map[string]*list.Element
}

type KeyPair struct {
	key   string
	value string
}

func New(capacity int) *Cache {
	l := &Cache{
		capacity: capacity,
		list:     new(list.List),
		elements: make(map[string]*list.Element, capacity),
	}
	return l
}
func (c *Cache) Get(key string) (string, error) {
	if val, ok := c.elements[key]; ok {
		c.list.MoveToFront(val)
		return val.Value.(*list.Element).Value.(*KeyPair).value, nil
	}
	return "", fmt.Errorf("not found")
}

func (c *Cache) Put(key, value string) {
	if node, ok := c.elements[key]; ok {
		c.list.MoveToFront(node)
		node.Value.(*list.Element).Value.(*KeyPair).value = value
		return
	} else if c.list.Len() == c.capacity {
		key := c.list.Back().Value.(*list.Element).Value.(*KeyPair).key
		delete(c.elements, key)
		c.list.Remove(c.list.Back())
	}
	node := &list.Element{Value: &KeyPair{
		key:   key,
		value: value,
	}}
	pointer := c.list.PushFront(node)
	c.elements[key] = pointer
}

func (cache *Cache) Print() {
	first := cache.list.Front()
	for first != nil {
		fmt.Printf("%s: %s\n", first.Value.(*list.Element).Value.(*KeyPair).key, first.Value.(*list.Element).Value.(*KeyPair).value)
		first = first.Next()
	}
}

func PlayCache() {
	c := New(10)
	c.Put("1", "a")
	c.Put("2", "b")
	c.Put("3", "c")
	c.Put("4", "d")
	c.Get("1")
	c.Put("2", "z")
	c.Print()
}
